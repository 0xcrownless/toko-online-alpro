package main

import "fmt"

const MAXKERANJANG = 100

type keranjang struct {
	namapembeli string
	namabarang  string
	jumlah      int
	total       int
}

var datakeranjang [MAXKERANJANG]keranjang
var jumlahkeranjang int

func tambahkeranjang() {
	var cart keranjang

	fmt.Println("========== TAMBAH KE KERANJANG ==========")
	if jumlahkeranjang >= MAXKERANJANG {
		fmt.Println("Keranjang sudah penuh")
		return
	}

	cart.namapembeli = namapembeli
	fmt.Print("Nama Barang : ")
	fmt.Scan(&cart.namabarang)
	fmt.Print("Jumlah : ")
	fmt.Scan(&cart.jumlah)

	if cart.jumlah <= 0 {
		fmt.Println("Jumlah barang harus lebih dari 0")
		return
	}
	indexBarang := cariindexbarang(0, cart.namabarang)
	if indexBarang == -1 {
		fmt.Println("Barang tidak ditemukan")
		return
	}

	for i := 0; i < jumlahkeranjang; i++ {
		if datakeranjang[i].namabarang == cart.namabarang {
			jumlahBaru := datakeranjang[i].jumlah + cart.jumlah
			if jumlahBaru > databarang[indexBarang].stok {
				fmt.Println("Stok", cart.namabarang, "tidak cukup")
				return
			}
			datakeranjang[i].jumlah = jumlahBaru
			fmt.Println("Jumlah barang di keranjang berhasil ditambahkan")
			return
		}
	}

	if cart.jumlah > databarang[indexBarang].stok {
		fmt.Println("Stok", cart.namabarang, "tidak cukup")
		return
	}
	datakeranjang[jumlahkeranjang] = cart
	jumlahkeranjang++
	fmt.Println("Barang berhasil masuk keranjang")
}

func tampilkeranjang() {
	fmt.Println("========== KERANJANG ==========")
	if jumlahkeranjang == 0 {
		fmt.Println("Keranjang kosong")
		return
	}

	for i := 0; i < jumlahkeranjang; i++ {
		fmt.Println("Pembeli :", datakeranjang[i].namapembeli)
		fmt.Println("Barang  :", datakeranjang[i].namabarang)
		fmt.Println("Jumlah  :", datakeranjang[i].jumlah)
		fmt.Println("------------------------")
	}

	var pilihan int
	fmt.Println("1. Checkout")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		checkoutkeranjang()
	}
}

func checkoutkeranjang() {
	if jumlahkeranjang == 0 {
		fmt.Println("Keranjang kosong")
		return
	}
	if jumlahtransaksi >= MAXTRANSAKSI {
		fmt.Println("Data transaksi sudah penuh")
		return
	}

	fmt.Println("========== CHECKOUT KERANJANG ==========")
	tampilpilihanpembayaran()
	var pilihanbayar int
	fmt.Print("Pilih pembayaran : ")
	fmt.Scan(&pilihanbayar)
	metode, valid := pilihmetodepembayaran(pilihanbayar)
	if !valid {
		fmt.Println("Metode pembayaran tidak tersedia")
		return
	}

	trx := transaksi{
		idtransaksi:      idtransaksibaru(),
		pembeli:          namapembeli,
		status:           "pending",
		metodepembayaran: metode,
	}
	for i := 0; i < jumlahkeranjang; i++ {
		index := cariindexbarang(0, datakeranjang[i].namabarang)
		if index == -1 {
			fmt.Println("Barang", datakeranjang[i].namabarang, "tidak ditemukan")
			return
		}
		if datakeranjang[i].jumlah <= 0 {
			fmt.Println("Jumlah barang harus lebih dari 0")
			return
		}
		trx.total += datakeranjang[i].jumlah * databarang[index].harga
		trx.jumlah += datakeranjang[i].jumlah
		if trx.namabarang != "" {
			trx.namabarang += ", "
		}
		trx.namabarang += databarang[index].nama + " x" + fmt.Sprint(datakeranjang[i].jumlah)
	}

	if !prosesstoktransaksi(trx, false) {
		return
	}
	datatransaksi[jumlahtransaksi] = trx
	jumlahtransaksi++
	savetransaksi()
	jumlahkeranjang = 0
	datakeranjang = [MAXKERANJANG]keranjang{}
	fmt.Println("Checkout berhasil")
}

func hapuskeranjang(namabarang string) {
	for i := 0; i < jumlahkeranjang; i++ {
		if datakeranjang[i].namabarang != namabarang {
			continue
		}
		for j := i; j < jumlahkeranjang-1; j++ {
			datakeranjang[j] = datakeranjang[j+1]
		}
		jumlahkeranjang--
		datakeranjang[jumlahkeranjang] = keranjang{}
		fmt.Println("Barang dihapus dari keranjang")
		return
	}
	fmt.Println("Barang tidak ditemukan")
}

func kosongkankeranjang() {
	jumlahkeranjang = 0
	datakeranjang = [MAXKERANJANG]keranjang{}
}
