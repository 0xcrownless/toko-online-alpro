package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAXADMIN = 50

type admin struct {
	usn  string
	pw   string
	role string
}

var dataadmin [MAXADMIN]admin
var jumlahadmin int
var adminlogin string

func login() string {
	var usn string
	var pw string

	fmt.Println("========== LOGIN ==========")
	fmt.Print("Username : ")
	fmt.Scan(&usn)
	fmt.Print("Password : ")
	fmt.Scan(&pw)

	for i := 0; i < jumlahadmin; i++ {
		if usn == dataadmin[i].usn && pw == dataadmin[i].pw {
			adminlogin = dataadmin[i].usn
			fmt.Println("Login berhasil")
			return dataadmin[i].role
		}
	}

	fmt.Println("Login gagal")
	return ""
}

func tambahadmin() {
	var adm admin

	fmt.Println("========== TAMBAH ADMIN ==========")
	if jumlahadmin >= MAXADMIN {
		fmt.Println("Data admin sudah penuh")
		return
	}

	fmt.Print("Username : ")
	fmt.Scan(&adm.usn)
	fmt.Print("Password : ")
	fmt.Scan(&adm.pw)

	if adm.usn == "" || adm.pw == "" {
		fmt.Println("Username dan password tidak boleh kosong")
		return
	}

	for i := 0; i < jumlahadmin; i++ {
		if dataadmin[i].usn == adm.usn {
			fmt.Println("Username sudah digunakan")
			return
		}
	}

	adm.role = "admin"
	dataadmin[jumlahadmin] = adm
	jumlahadmin++
	saveadmin()
	fmt.Println("Admin berhasil ditambahkan")
}

func tampiladmin() {
	fmt.Println("========== DATA ADMIN ==========")
	for i := 0; i < jumlahadmin; i++ {
		fmt.Println("Username :", dataadmin[i].usn)
		fmt.Println("Role     :", dataadmin[i].role)
		fmt.Println("----------------------")
	}
}

func saveadmin() {
	file, err := os.Create("admin.txt")
	if err != nil {
		fmt.Println("Gagal menyimpan data admin")
		return
	}
	defer file.Close()

	for i := 0; i < jumlahadmin; i++ {
		data := dataadmin[i].usn + "|" + dataadmin[i].pw + "|" + dataadmin[i].role + "\n"
		if _, err := file.WriteString(data); err != nil {
			fmt.Println("Gagal menyimpan data admin")
			return
		}
	}
}

func loadadmin() {
	file, err := os.Open("admin.txt")
	if err != nil {
		return
	}
	defer file.Close()

	jumlahadmin = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "|")
		if len(data) < 3 || jumlahadmin >= MAXADMIN {
			continue
		}
		if data[0] == "" || data[1] == "" || (data[2] != "admin" && data[2] != "superadmin") {
			continue
		}
		dataadmin[jumlahadmin] = admin{usn: data[0], pw: data[1], role: data[2]}
		jumlahadmin++
	}
}

func hapusadmin() {
	var usn string

	fmt.Println("========== HAPUS ADMIN ==========")
	fmt.Print("Masukkan username admin : ")
	fmt.Scan(&usn)

	for i := 0; i < jumlahadmin; i++ {
		if dataadmin[i].usn != usn {
			continue
		}
		if dataadmin[i].role == "superadmin" {
			fmt.Println("Superadmin tidak dapat dihapus")
			return
		}
		for j := i; j < jumlahadmin-1; j++ {
			dataadmin[j] = dataadmin[j+1]
		}
		jumlahadmin--
		dataadmin[jumlahadmin] = admin{}
		saveadmin()
		fmt.Println("Admin berhasil dihapus")
		return
	}

	fmt.Println("Admin tidak ditemukan")
}
