package main

import "fmt"

func menuuser() {

	var pilihan int

	for {

		fmt.Println("========== MENU PEMBELI ==========")
		fmt.Println("1. Lihat Barang")
		fmt.Println("2. Cari Barang")
		fmt.Println("3. Beli Barang")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			tampilbarang()

		} else if pilihan == 2 {

			caribarang()

		} else if pilihan == 3 {

			tambahtransaksi()

		} else if pilihan == 0 {

			fmt.Println("Keluar dari menu pembeli")
			break

		} else {

			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}