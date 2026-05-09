package main

import "fmt"

func selectionsorthargaasc() {

	var i int
	var j int
	var min int
	var temp barang

	for i = 0; i < jumlahbarang-1; i++ {

		min = i

		for j = i + 1; j < jumlahbarang; j++ {

			if databarang[j].harga < databarang[min].harga {

				min = j
			}
		}

		temp = databarang[i]
		databarang[i] = databarang[min]
		databarang[min] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan harga ascending")
}

func insertionsortnamaasc() {

	var i int
	var j int
	var temp barang

	for i = 1; i < jumlahbarang; i++ {

		temp = databarang[i]
		j = i - 1

		for j >= 0 && databarang[j].nama > temp.nama {

			databarang[j+1] = databarang[j]
			j--
		}

		databarang[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan nama ascending")
}

func menusorting() {

	var pilihan int

	for {

		fmt.Println("========== MENU SORTING ==========")
		fmt.Println("1. Harga Ascending")
		fmt.Println("2. Nama Ascending")
		fmt.Println("0. Kembali")

		fmt.Print("Pilih : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			selectionsorthargaasc()
			tampilbarang()

		} else if pilihan == 2 {

			insertionsortnamaasc()
			tampilbarang()

		} else if pilihan == 0 {

			break

		} else {

			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}