package main

import (
	"fmt"
	"os"
	"bufio"
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

func initadmin() {

	dataadmin[0].usn = "reval"
	dataadmin[0].pw = "123"
	dataadmin[0].role = "superadmin"

	jumlahadmin = 1
}

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

	var admin admin

	fmt.Println("========== TAMBAH ADMIN ==========")

	fmt.Print("Username : ")
	fmt.Scan(&admin.usn)

	fmt.Print("Password : ")
	fmt.Scan(&admin.pw)

	admin.role = "admin"

	dataadmin[jumlahadmin] = admin
	jumlahadmin++

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
	var admin admin

	file, _ = os.Open("admin.txt")

	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {

		line = scanner.Text()

		data = strings.Split(line, "|")

		admin.usn = data[0]
		admin.pw = data[1]
		admin.role = data[2]

		dataadmin[jumlahadmin] = admin
		jumlahadmin++
	}
}