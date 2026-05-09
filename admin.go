package main 
import "fmt"
const MAXADMIN = 50
type admin struct {
	usn string 
	pw string 
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

	fmt.Println("========== login ==========")
	fmt.Print("usn: ")
	fmt.Scan(&usn)
	fmt.Print("password : ")
	fmt.Scan(&pw)

	for i = 0; i < jumlahadmin; i++ {
		if usn == dataadmin[i].usn && pw == dataadmin[i].pw {
			role = dataadmin[i].role 
		}
	}
	if role == "" {
	fmt.Println("login gagal")
	} else {
		fmt.Println("login berhasil")
	}
	return role 
}

func tambahadmin() {
	var admin admin 

	fmt.Println("========== Tambah admin ==========")

	fmt.Print("username: ")
	fmt.Scan(&admin.usn)
	fmt.Print("password: ")
	fmt.Scan(&admin.pw)

	admin.role = "admin"

	dataadmin[jumlahadmin] = admin
	jumlahadmin++

	fmt.Println("admin berhasil ditambakan")
}

func tampiladmin() {
	var i int 
	fmt.Println("========== DATA ADMIN ==========")
	for i = 0; i < jumlahadmin; i++ {
		
		fmt.Println("username:", dataadmin[i].usn)
		fmt.Println("role:", dataadmin[i].role)
		fmt.Println("")
	}
}