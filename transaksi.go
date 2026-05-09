package main

import import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const MAXTRANSAKSI = 999

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
		savetransaksi()

	} else {

		fmt.Println("Barang tidak ditemukan")
	}
}

func approvetransaksi() {

	var idtrx int
	var i int
	var j int
	var ketemu bool

	ketemu = false

	fmt.Println("========== APPROVE TRANSAKSI ==========")

	fmt.Print("Masukkan ID Transaksi : ")
	fmt.Scan(&idtrx)

	for i = 0; i < jumlahtransaksi; i++ {

		if datatransaksi[i].idtransaksi == idtrx {

			ketemu = true

			if datatransaksi[i].status == "approved" {

				fmt.Println("Transaksi sudah diapprove")

			} else {

				datatransaksi[i].status = "approved"

				for j = 0; j < jumlahbarang; j++ {

					if databarang[j].id == datatransaksi[i].idbarang {

						databarang[j].stok =
							databarang[j].stok - datatransaksi[i].jumlah

						databarang[j].terjual =
							databarang[j].terjual + datatransaksi[i].jumlah
					}
				}

				fmt.Println("Transaksi berhasil diapprove")
				savetransaksi()
				savebarang()
			}
		}
	}

	if ketemu == false {

		fmt.Println("Transaksi tidak ditemukan")
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

func savetransaksi() {

	var file *os.File
	var data string
	var i int

	file, _ = os.Create("transaksi.txt")

	defer file.Close()

	for i = 0; i < jumlahtransaksi; i++ {

		data =
			strconv.Itoa(datatransaksi[i].idtransaksi) + "|" +
				datatransaksi[i].pembeli + "|" +
				strconv.Itoa(datatransaksi[i].idbarang) + "|" +
				datatransaksi[i].namabarang + "|" +
				strconv.Itoa(datatransaksi[i].jumlah) + "|" +
				strconv.Itoa(datatransaksi[i].total) + "|" +
				datatransaksi[i].status + "\n"

		file.WriteString(data)
	}
}

func loadtransaksi() {

	var file *os.File
	var scanner *bufio.Scanner
	var line string
	var data []string
	var trx transaksi

	file, _ = os.Open("transaksi.txt")

	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {

		line = scanner.Text()

		data = strings.Split(line, "|")

		trx.idtransaksi, _ =
			strconv.Atoi(data[0])

		trx.pembeli = data[1]

		trx.idbarang, _ =
			strconv.Atoi(data[2])

		trx.namabarang = data[3]

		trx.jumlah, _ =
			strconv.Atoi(data[4])

		trx.total, _ =
			strconv.Atoi(data[5])

		trx.status = data[6]

		datatransaksi[jumlahtransaksi] = trx
		jumlahtransaksi++
	}
}