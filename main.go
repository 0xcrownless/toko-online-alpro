package main

import "fmt"

func main() {

	var role string
	var pilihan int

	initadmin()

	role = login()

	fmt.Println("Role :", role)

	if role == "superadmin" {

		for pilihan != 0 {

			fmt.Println("===== SUPER ADMIN AKSES =====")
			fmt.Println("1. Tambah Admin")
			fmt.Println("2. Lihat Admin")
			fmt.Println("3. Masuk Sistem Toko")
			fmt.Println("0. Keluar")

			fmt.Print("Pilih menu : ")
			fmt.Scan(&pilihan)

			if pilihan == 1 {

				tambahadmin()

			} else if pilihan == 2 {

				tampiladmin()

			} else if pilihan == 3 {

				menutoko()

			} else if pilihan == 0 {

				fmt.Println("Program selesai")

			} else {

				fmt.Println("Menu tidak tersedia")
			}

			fmt.Println()
		}

	} else if role == "admin" {

		menutoko()

	} else {

		fmt.Println("Role tidak dikenali")
	}
}

func menutoko() {

	var pilihan int

	for pilihan != 0 {

		fmt.Println("========== TOKO ONLINE ==========")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			tambahbarang()

		} else if pilihan == 2 {

			tampilbarang()

		} else if pilihan == 3 {

			editbarang()

		} else if pilihan == 4 {

			hapusbarang()

		} else if pilihan == 0 {

			fmt.Println("Kembali")

		} else {

			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}