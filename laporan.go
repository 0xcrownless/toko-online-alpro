package main

import "fmt"

func topbarangterlaris() {
	var peringkat [MAXBARANG]barang
	for i := 0; i < jumlahbarang; i++ {
		peringkat[i] = databarang[i]
	}

	for i := 0; i < jumlahbarang-1; i++ {
		for j := i + 1; j < jumlahbarang; j++ {
			if peringkat[j].terjual > peringkat[i].terjual {
				peringkat[i], peringkat[j] = peringkat[j], peringkat[i]
			}
		}
	}

	fmt.Println("========== TOP BARANG TERLARIS ==========")
	for i := 0; i < jumlahbarang && i < 3; i++ {
		fmt.Println("Nama     :", peringkat[i].nama)
		fmt.Println("Terjual  :", peringkat[i].terjual)
		fmt.Println("-------------------------")
	}
}
