package main

import "fmt"

var namapembeli string

func menuuser() {
	var pilihan int

	kosongkankeranjang()
	suara("welcome, and happy shopping")
	fmt.Print("nama pembeli : ")
	fmt.Scan(&namapembeli)

	for {
		fmt.Println("========== MENU PEMBELI ==========")
		fmt.Println("1. Lihat Barang")
		fmt.Println("2. Cari Barang")
		fmt.Println("3. Beli Barang")
		fmt.Println("4. Tambah ke Keranjang")
		fmt.Println("5. Lihat Keranjang")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilbarang()
		case 2:
			caribarang()
		case 3:
			tambahtransaksi()
		case 4:
			tambahkeranjang()
		case 5:
			tampilkeranjang()
		case 0:
			kosongkankeranjang()
			fmt.Println("Keluar dari menu pembeli")
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
		fmt.Println()
	}
}
