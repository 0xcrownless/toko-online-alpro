package main

import "fmt"

func stastistiktoko() {
	var totalpendapatan int
	var totaltransaksi int
	var totalbarang int

	for i := 0; i < jumlahtransaksi; i++ {
		if datatransaksi[i].status == "approved" {
			totalpendapatan += datatransaksi[i].total
			totaltransaksi++
			totalbarang += datatransaksi[i].jumlah
		}
	}

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║           STATISTIK TOKO             ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Printf("║ %-20s : %-11d   ║\n", "Total Penjualan", totalpendapatan)
	fmt.Printf("║ %-20s : %-11d   ║\n", "Total Penarikan", totalpenarikan)
	fmt.Printf("║ %-20s : %-11d   ║\n", "Saldo Toko", saldotoko)
	fmt.Printf("║ %-20s : %-11d   ║\n", "Total Transaksi", totaltransaksi)
	fmt.Printf("║ %-20s : %-11d   ║\n", "Barang Terjual", totalbarang)
	fmt.Println("╚══════════════════════════════════════╝")
}
