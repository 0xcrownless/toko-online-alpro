package main

import "fmt"

func main() {

	var pilihan int
	var role string

	loadadmin()
	loadbarang()
	loadtransaksi()
	loadpenarikan()

	for {

		fmt.Println("========== TOKO ONLINE ==========")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Belanja")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			role = login()

			fmt.Println("Role :", role)

			if role == "superadmin" {

				suara("welcome, developer, Reval ")

				for {

					fmt.Println("===== SUPER ADMIN AKSES =====")
					fmt.Println("1. Tambah Admin")
					fmt.Println("2. Lihat Admin")
					fmt.Println("3. hapus admin")
					fmt.Println("4. Masuk Sistem Toko")
					fmt.Println("0. Keluar")

					fmt.Print("Pilih menu : ")
					fmt.Scan(&pilihan)

					if pilihan == 1 {

						tambahadmin()

					} else if pilihan == 2 {

						tampiladmin()

					} else if pilihan == 3 {

						hapusadmin()

					} else if pilihan == 4 {

						menutoko(role)

					} else if pilihan == 0 {

						break

					} else {

						fmt.Println("Menu tidak tersedia")
					}

					fmt.Println()
				}

			} else if role == "admin" {

				suara("welcome, " + adminlogin)

				menutoko(role)

			}

		} else if pilihan == 2 {

			menuuser()

		} else if pilihan == 0 {

			fmt.Println("Program selesai")
			break

		} else {

			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}

func menutoko(role string) {

	var pilihan int

	for {

		fmt.Println("========== TOKO ONLINE ==========")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Sorting Barang")
		fmt.Println("7. Tambah Transaksi")
		fmt.Println("8. Tampilkan Transaksi")
		fmt.Println("9. Approve Transaksi")
		fmt.Println("10. hapus transaksi")

		if role == "superadmin" {

			fmt.Println("11. Top Barang Terlaris")
			fmt.Println("12. Export Excel")
			fmt.Println("13. statistik  toko")
			fmt.Println("14. tarik saldo  ")
			fmt.Println("15. Tutup Laporan Bulanan")
		}

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

		} else if pilihan == 5 {

			caribarang()

		} else if pilihan == 6 {

			menusorting()

		} else if pilihan == 7 {

			tambahtransaksiadmin()

		} else if pilihan == 8 {

			tampiltransaksi()

		} else if pilihan == 9 {

			approvetransaksi()

		} else if pilihan == 10 {

			hapustransaksi()

		} else if pilihan == 11 && role == "superadmin" {

			topbarangterlaris()

		} else if pilihan == 12 && role == "superadmin" {

			exportexcel()

		} else if pilihan == 13 && role == "superadmin" {

			stastistiktoko()

		} else if pilihan == 14 && role == "superadmin" {

			tariksaldo()

		} else if pilihan == 15 && role == "superadmin" {

			tutuplaporanbulanan()

		} else if pilihan == 0 {

			fmt.Println("Kembali")
			break

		} else {

			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}
