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

	var nama string
	var index int

	fmt.Println("========== CARI BARANG ==========")

	fmt.Print("Masukkan nama barang : ")
	fmt.Scan(&nama)

	index = sequentialsearchnama(nama)

	if index == -1 {

		fmt.Println("Barang tidak ditemukan")

	} else {

		fmt.Println("Barang ditemukan")
		fmt.Println("ID        :", databarang[index].id)
		fmt.Println("Nama      :", databarang[index].nama)
		fmt.Println("Kategori  :", databarang[index].kategori)
		fmt.Println("Harga     :", databarang[index].harga)
		fmt.Println("Stok      :", databarang[index].stok)
		fmt.Println("Terjual   :", databarang[index].terjual)
	}
}