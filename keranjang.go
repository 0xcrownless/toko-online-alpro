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

	cart.namapembeli = namapembeli

	fmt.Print("Nama Barang : ")
	fmt.Scan(&cart.namabarang)

	fmt.Print("Jumlah : ")
	fmt.Scan(&cart.jumlah)

	datakeranjang[jumlahkeranjang] = cart
	jumlahkeranjang++

	fmt.Println("Barang berhasil masuk keranjang")
}

func tampilkeranjang() {

	var i int
	var pilihan int

	fmt.Println("========== KERANJANG ==========")

	if jumlahkeranjang == 0 {

		fmt.Println("Keranjang kosong")
		return
	}

	for i = 0; i < jumlahkeranjang; i++ {

		fmt.Println("Pembeli :", datakeranjang[i].namapembeli)
		fmt.Println("Barang  :", datakeranjang[i].namabarang)
		fmt.Println("Jumlah  :", datakeranjang[i].jumlah)
		fmt.Println("------------------------")
	}

	fmt.Println("1. Checkout")
	fmt.Println("0. Kembali")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {

		checkoutkeranjang()
	}
}

func checkoutkeranjang() {

	var i int
	var j int
	var trx transaksi
	var pilihanbayar int
	var semuabarang string

	if jumlahkeranjang == 0 {

		fmt.Println("Keranjang kosong")
		return
	}

	fmt.Println("========== CHECKOUT KERANJANG ==========")

	fmt.Println("1. Dana")
	fmt.Println("2. Ovo")
	fmt.Println("3. Gopay")
	fmt.Println("4. QRIS")
	fmt.Println("5. Bank BCA")
	fmt.Println("6. Bank BRI")
	fmt.Println("7. Bank Mandiri")

	fmt.Print("Pilih pembayaran : ")
	fmt.Scan(&pilihanbayar)

	if jumlahtransaksi == 0 {

		trx.idtransaksi = 1

	} else {

		trx.idtransaksi =
			datatransaksi[jumlahtransaksi-1].idtransaksi + 1
	}

	trx.pembeli = namapembeli
	trx.status = "pending"

	for i = 0; i < jumlahkeranjang; i++ {

		for j = 0; j < jumlahbarang; j++ {

			if datakeranjang[i].namabarang ==
				databarang[j].nama {

				if datakeranjang[i].jumlah >
					databarang[j].stok {
					
					fmt.Println("-------------------------------------------")
					fmt.Println("Stok", databarang[j].nama, "tidak cukup")
					
					fmt.Println("1. hapus dari kerangjang ")
					fmt.Println("0. batal")

					var pilih int 

					fmt.Print("pilih : ")
					fmt.Scan(&pilih)

					if pilih == 1 {

						hapuskeranjang(datakeranjang[i].namabarang)
					}
					return
				}

				trx.total +=
					datakeranjang[i].jumlah *
						databarang[j].harga

				semuabarang +=
					databarang[j].nama +
						" x" +
						fmt.Sprint(datakeranjang[i].jumlah)

				if i != jumlahkeranjang-1 {

					semuabarang += ", "
				}
			}
		}
	}

	trx.namabarang = semuabarang

	if pilihanbayar == 1 {

		trx.metodepembayaran = "Dana"

	} else if pilihanbayar == 2 {

		trx.metodepembayaran = "Ovo"

	} else if pilihanbayar == 3 {

		trx.metodepembayaran = "Gopay"

	} else if pilihanbayar == 4 {

		trx.metodepembayaran = "QRIS"

	} else if pilihanbayar == 5 {

		trx.metodepembayaran = "Bank BCA"

	} else if pilihanbayar == 6 {

		trx.metodepembayaran = "Bank BRI"

	} else if pilihanbayar == 7 {

		trx.metodepembayaran = "Bank Mandiri"
	}

	datatransaksi[jumlahtransaksi] = trx
	jumlahtransaksi++

	savetransaksi()

	jumlahkeranjang = 0

	fmt.Println("Checkout berhasil")
}

func hapuskeranjang(namabarang string) {

	var i int
	var ketemu bool

	ketemu = false

	for i = 0; i < jumlahkeranjang; i++ {

		if datakeranjang[i].namabarang == namabarang {

			ketemu = true

			for i < jumlahkeranjang-1 {

				datakeranjang[i] =
					datakeranjang[i+1]

				i++
			}

			jumlahkeranjang--

			fmt.Println("Barang dihapus dari keranjang")
		}
	}

	if ketemu == false {

		fmt.Println("Barang tidak ditemukan")
	}
}