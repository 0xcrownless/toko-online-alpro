package main

import "fmt"

func sequentialsearchnama(nama string) int {

	var i int
	var index int

	index = -1

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].nama == nama {

			index = i
		}
	}

	return index
}

func caribarang() {

	var keyword string
	var i int
	var ketemu bool

	ketemu = false

	fmt.Println("========== CARI BARANG ==========")

	fmt.Print("Masukkan nama / kategori : ")
	fmt.Scan(&keyword)

	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        HASIL PENCARIAN                           ║")
	fmt.Println("╠════╦════════════════════╦══════════════╦════════════╦════════════╣")
	fmt.Println("║ ID ║ Nama Barang        ║ Kategori     ║ Harga      ║ Stok       ║")
	fmt.Println("╠════╬════════════════════╬══════════════╬════════════╬════════════╣")

	for i = 0; i < jumlahbarang; i++ {

		if keyword == databarang[i].nama ||
			keyword == databarang[i].kategori {

			ketemu = true

			fmt.Printf(
				"║ %-2d ║ %-18s ║ %-12s ║ %-10d ║ %-10d ║\n",
				databarang[i].id,
				databarang[i].nama,
				databarang[i].kategori,
				databarang[i].harga,
				databarang[i].stok,
			)
		}
	}

	fmt.Println("╚════╩════════════════════╩══════════════╩════════════╩════════════╝")

	if ketemu == false {

		fmt.Println("Barang tidak ditemukan")
	}
}
