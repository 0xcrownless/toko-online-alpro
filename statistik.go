package main

import "fmt"

func stastistiktoko() {

	var i int

	var omzet int
	var totaltransaksi int
	var totalbarang int

	for i = 0; i < jumlahtransaksi; i++ {

		if datatransaksi[i].status == "approved" {

			omzet += datatransaksi[i].total

			totaltransaksi++

			totalbarang += datatransaksi[i].jumlah
		}
	}

	fmt.Println("========== DASHBOARD TOKO ==========")

	fmt.Println("Saldo Toko      : Rp", saldotoko)
	fmt.Println("Omzet Toko      : Rp", omzet)
	fmt.Println("Total Transaksi :", totaltransaksi)
	fmt.Println("Barang Terjual  :", totalbarang)
}