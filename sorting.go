package main

import "fmt"

func selectionsorthargaasc() {
	for i := 0; i < jumlahbarang-1; i++ {
		min := i
		for j := i + 1; j < jumlahbarang; j++ {
			if databarang[j].harga < databarang[min].harga {
				min = j
			}
		}
		databarang[i], databarang[min] = databarang[min], databarang[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan harga ascending")
}

func insertionsortnamaasc() {
	for i := 1; i < jumlahbarang; i++ {
		temp := databarang[i]
		j := i - 1
		for j >= 0 && databarang[j].nama > temp.nama {
			databarang[j+1] = databarang[j]
			j--
		}
		databarang[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama ascending")
}

func menusorting() {
	for {
		var pilihan int
		fmt.Println("========== MENU SORTING ==========")
		fmt.Println("1. Harga Ascending")
		fmt.Println("2. Nama Ascending")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			selectionsorthargaasc()
			tampilbarangsaatini()
		case 2:
			insertionsortnamaasc()
			tampilbarangsaatini()
		case 0:
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
		fmt.Println()
	}
}
