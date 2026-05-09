package main

import "fmt"

func laporanpenjualan() {

	var i int
	var totalomzet int
	var totaltransaksi int

	totalomzet = 0
	totaltransaksi = 0

	fmt.Println("========== LAPORAN PENJUALAN ==========")

	for i = 0; i < jumlahtransaksi; i++ {

		if datatransaksi[i].status == "approved" {

			totalomzet =
				totalomzet + datatransaksi[i].total

			totaltransaksi++
		}
	}

	fmt.Println("Total Transaksi :", totaltransaksi)
	fmt.Println("Total Omzet     :", totalomzet)
}

func topbarangterlaris() {

	var i int
	var j int
	var temp barang

	for i = 0; i < jumlahbarang-1; i++ {

		for j = i + 1; j < jumlahbarang; j++ {

			if databarang[j].terjual >
				databarang[i].terjual {

				temp = databarang[i]
				databarang[i] = databarang[j]
				databarang[j] = temp
			}
		}
	}

	fmt.Println("========== TOP BARANG TERLARIS ==========")

	for i = 0; i < jumlahbarang && i < 3; i++ {

		fmt.Println("Nama     :", databarang[i].nama)
		fmt.Println("Terjual  :", databarang[i].terjual)
		fmt.Println("-------------------------")
	}
}