package main

import  (
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
	metodepembayaran string
}

var datatransaksi [MAXTRANSAKSI]transaksi
var jumlahtransaksi int

func tambahtransaksi() {

	var trx transaksi
	var namabarang string
	var i int
	var ketemu bool
	var pilihanbayar int 

	ketemu = false

	if jumlahtransaksi == 0 {

		trx.idtransaksi = 1

	} else {

		trx.idtransaksi =
			datatransaksi[jumlahtransaksi-1].idtransaksi + 1
	}

	fmt.Println("========== TRANSAKSI PEMBELIAN ==========")

	fmt.Print("Nama Pembeli : ")
	fmt.Scan(&trx.pembeli)

	fmt.Print("Nama Barang : ")
	fmt.Scan(&namabarang)

	fmt.Print("Jumlah Beli : ")
	fmt.Scan(&trx.jumlah)

	fmt.Println("========== PILIH PEMBAYARAN ==========")
	fmt.Println("1. Dana")
	fmt.Println("2. Ovo")
	fmt.Println("3. Gopay")
	fmt.Println("4. QRIS")
	fmt.Println("5. Bank BCA")
	fmt.Println("6. Bank BRI")
	fmt.Println("7. Bank Mandiri")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilihanbayar)

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

	} else {

		fmt.Println("Metode pembayaran tidak tersedia")
		return
}

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].nama == namabarang {

			ketemu = true

			if trx.jumlah > databarang[i].stok {

				fmt.Println("Stok tidak cukup")
				return
			}

			trx.idbarang = databarang[i].id
			trx.namabarang = databarang[i].nama

			trx.total =
				trx.jumlah * databarang[i].harga

			trx.status = "pending"

			datatransaksi[jumlahtransaksi] = trx
			jumlahtransaksi++

			savetransaksi()

			fmt.Println("Transaksi berhasil ditambahkan")
		}
	}

	if ketemu == false {

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

				saldotoko += datatransaksi[i].total
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

func hapustransaksi() {

	var id int
	var i int
	var j int
	var ketemu bool

	ketemu = false

	fmt.Println("========== HAPUS TRANSAKSI ==========")

	fmt.Print("Masukkan ID transaksi : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahtransaksi; i++ {

		if datatransaksi[i].idtransaksi == id {

			ketemu = true

			for j = i; j < jumlahtransaksi-1; j++ {

				datatransaksi[j] = datatransaksi[j+1]
			}

			jumlahtransaksi--

			savetransaksi()

			fmt.Println("Transaksi berhasil dihapus")
		}
	}

	if ketemu == false {

		fmt.Println("Transaksi tidak ditemukan")
	}
}