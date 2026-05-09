package main

import "fmt"

const MAXBARANG = 100

type barang struct {
	id        int
	nama      string
	kategori  string
	harga     int
	stok      int
	terjual   int
}

var databarang [MAXBARANG]barang
var jumlahbarang int

func tambahbarang() {

	var barang barang

	fmt.Println("========== TAMBAH BARANG ==========")

	fmt.Print("ID Barang : ")
	fmt.Scan(&barang.id)

	fmt.Print("Nama Barang : ")
	fmt.Scan(&barang.nama)

	fmt.Print("Kategori : ")
	fmt.Scan(&barang.kategori)

	fmt.Print("Harga : ")
	fmt.Scan(&barang.harga)

	fmt.Print("Stok : ")
	fmt.Scan(&barang.stok)

	barang.terjual = 0

	databarang[jumlahbarang] = barang
	jumlahbarang++

	fmt.Println("Barang berhasil ditambahkan")
}

func tampilbarang() {

	var i int

	fmt.Println("========== DATA BARANG ==========")

	if jumlahbarang == 0 {

		fmt.Println("Data barang kosong")

	} else {

		for i = 0; i < jumlahbarang; i++ {

			fmt.Println("ID        :", databarang[i].id)
			fmt.Println("Nama      :", databarang[i].nama)
			fmt.Println("Kategori  :", databarang[i].kategori)
			fmt.Println("Harga     :", databarang[i].harga)
			fmt.Println("Stok      :", databarang[i].stok)
			fmt.Println("Terjual   :", databarang[i].terjual)
			fmt.Println("-----------------------------")
		}
	}
}

func editbarang() {

	var id int
	var i int
	var pilihan int
	var ketemu bool

	ketemu = false

	fmt.Println("========== EDIT BARANG ==========")

	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].id == id {

			ketemu = true

			for {

				fmt.Println("===== PILIH DATA YANG DIUBAH =====")
				fmt.Println("1. Nama")
				fmt.Println("2. Kategori")
				fmt.Println("3. Harga")
				fmt.Println("4. Stok")
				fmt.Println("0. Selesai")

				fmt.Print("Pilih : ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {

					fmt.Print("Nama Baru : ")
					fmt.Scan(&databarang[i].nama)

				} else if pilihan == 2 {

					fmt.Print("Kategori Baru : ")
					fmt.Scan(&databarang[i].kategori)

				} else if pilihan == 3 {

					fmt.Print("Harga Baru : ")
					fmt.Scan(&databarang[i].harga)

				} else if pilihan == 4 {

					fmt.Print("Stok Baru : ")
					fmt.Scan(&databarang[i].stok)

				} else if pilihan == 0 {

					fmt.Println("Edit selesai")
					break

				} else {

					fmt.Println("Menu tidak tersedia")
				}

				fmt.Println()
			}
		}
	}

	if ketemu == false {

		fmt.Println("Barang tidak ditemukan")
	}
}

func hapusbarang() {

	var id int
	var i int
	var j int
	var ketemu bool

	ketemu = false

	fmt.Println("========== HAPUS BARANG ==========")

	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].id == id {

			ketemu = true

			for j = i; j < jumlahbarang-1; j++ {

				databarang[j] = databarang[j+1]
			}

			jumlahbarang--

			fmt.Println("Barang berhasil dihapus")
		}
	}

	if ketemu == false {

		fmt.Println("Barang tidak ditemukan")
	}
}