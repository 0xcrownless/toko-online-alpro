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
	var role string
	var i int

	role = ""

	fmt.Println("========== LOGIN ==========")

	fmt.Print("Username : ")
	fmt.Scan(&usn)

	fmt.Print("Password : ")
	fmt.Scan(&pw)

	for i = 0; i < jumlahadmin; i++ {

		if usn == dataadmin[i].usn &&
			pw == dataadmin[i].pw {

			role = dataadmin[i].role

			adminlogin = dataadmin[i].usn
		}
	}

	if role == "" {

		fmt.Println("Login gagal")

	} else {

		fmt.Println("Login berhasil")
	}

	return role
}

func tambahadmin() {

	var adm admin

	fmt.Println("========== TAMBAH ADMIN ==========")

	fmt.Print("Username : ")
	fmt.Scan(&adm.usn)

	fmt.Print("Password : ")
	fmt.Scan(&adm.pw)

	adm.role = "admin"

	dataadmin[jumlahadmin] = adm
	jumlahadmin++

	saveadmin()

	fmt.Println("Admin berhasil ditambahkan")
}

func tampiladmin() {

	var i int

	fmt.Println("========== DATA ADMIN ==========")

	for i = 0; i < jumlahadmin; i++ {

		fmt.Println("Username :", dataadmin[i].usn)
		fmt.Println("Role     :", dataadmin[i].role)
		fmt.Println("----------------------")
	}
}

func saveadmin() {

	var file *os.File
	var data string
	var i int

	file, _ = os.Create("admin.txt")

	defer file.Close()

	for i = 0; i < jumlahadmin; i++ {

		data =
			dataadmin[i].usn + "|" +
				dataadmin[i].pw + "|" +
				dataadmin[i].role + "\n"

		file.WriteString(data)
	}
}

func loadadmin() {

	var file *os.File
	var scanner *bufio.Scanner
	var line string
	var data []string
	var adm admin

	file, _ = os.Open("admin.txt")

	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {

		line = scanner.Text()

		data = strings.Split(line, "|")

		adm.usn = data[0]
		adm.pw = data[1]
		adm.role = data[2]

		dataadmin[jumlahadmin] = adm
		jumlahadmin++
	}
}

func hapusadmin() {

	var usn string
	var i int
	var j int
	var ketemu bool

	ketemu = false

	fmt.Println("========== HAPUS ADMIN ==========")

	fmt.Print("Masukkan username admin : ")
	fmt.Scan(&usn)

	for i = 0; i < jumlahadmin; i++ {

		if dataadmin[i].usn == usn {

			ketemu = true

			for j = i; j < jumlahadmin-1; j++ {

				dataadmin[j] = dataadmin[j+1]
			}

			jumlahadmin--

			fmt.Println("Admin berhasil dihapus")

			saveadmin()
		}
	}

	if ketemu == false {

		fmt.Println("Admin tidak ditemukan")
	}
}