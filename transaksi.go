package main

import "fmt"

const MAXTRANSAKSI = 100

type transaksi struct {
	idtransaksi int
	pembeli     string
	idbarang    int
	namabarang  string
	jumlah      int
	total       int
	status      string
}

var datatransaksi [MAXTRANSAKSI]transaksi
var jumlahtransaksi int

func tambahtransaksi() {

	var trx transaksi
	var i int
	var ketemu bool

	ketemu = false

	fmt.Println("========== TRANSAKSI PEMBELIAN ==========")

	fmt.Print("ID Transaksi : ")
	fmt.Scan(&trx.idtransaksi)

	fmt.Print("Nama Pembeli : ")
	fmt.Scan(&trx.pembeli)

	fmt.Print("ID Barang : ")
	fmt.Scan(&trx.idbarang)

	fmt.Print("Jumlah Beli : ")
	fmt.Scan(&trx.jumlah)

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].id == trx.idbarang {

			ketemu = true

			trx.namabarang = databarang[i].nama
			trx.total = trx.jumlah * databarang[i].harga
			trx.status = "pending"
		}
	}

	if ketemu == true {

		datatransaksi[jumlahtransaksi] = trx
		jumlahtransaksi++

		fmt.Println("Transaksi berhasil ditambahkan")
		fmt.Println("Status :", trx.status)
		fmt.Println("Total :", trx.total)

	} else {

		fmt.Println("Barang tidak ditemukan")
	}
}

func tampiltransaksi() {

	var i int

	fmt.Println("========== DATA TRANSAKSI ==========")

	for i = 0; i < jumlahtransaksi; i++ {

		fmt.Println("ID Transaksi :", datatransaksi[i].idtransaksi)
		fmt.Println("Pembeli      :", datatransaksi[i].pembeli)
		fmt.Println("Barang        :", datatransaksi[i].namabarang)
		fmt.Println("Jumlah        :", datatransaksi[i].jumlah)
		fmt.Println("Total         :", datatransaksi[i].total)
		fmt.Println("Status        :", datatransaksi[i].status)
		fmt.Println("-------------------------------")
	}
}