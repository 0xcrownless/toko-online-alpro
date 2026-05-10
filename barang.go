package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const MAXBARANG = 999

type barang struct {
	id        int
	nama      string
	kategori  string
	harga     int
	stok      int
	terjual   int
}

var databarang [MAXBARANG]barang
var jumlahbarang int

func tambahbarang() {

	var barang barang

	fmt.Println("========== TAMBAH BARANG ==========")

	barang.id = jumlahbarang + 1
	
	fmt.Print("Nama Barang : ")
	fmt.Scan(&barang.nama)

	fmt.Print("Kategori : ")
	fmt.Scan(&barang.kategori)

	fmt.Print("Harga : ")
	fmt.Scan(&barang.harga)

	fmt.Print("Stok : ")
	fmt.Scan(&barang.stok)

	barang.terjual = 0

	databarang[jumlahbarang] = barang
	jumlahbarang++

	fmt.Println("Barang berhasil ditambahkan")
	savebarang()
}

func tampilbarang() {

	var i int

	fmt.Println("========== DATA BARANG ==========")

	if jumlahbarang == 0 {

		fmt.Println("Data barang kosong")

	} else {

		for i = 0; i < jumlahbarang; i++ {

			fmt.Println("ID        :", databarang[i].id)
			fmt.Println("Nama      :", databarang[i].nama)
			fmt.Println("Kategori  :", databarang[i].kategori)
			fmt.Println("Harga     :", databarang[i].harga)
			fmt.Println("Stok      :", databarang[i].stok)
			fmt.Println("Terjual   :", databarang[i].terjual)
			fmt.Println("-----------------------------")
		}
	}
}

func editbarang() {

	var id int
	var i int
	var pilihan int
	var ketemu bool

	ketemu = false

	fmt.Println("========== EDIT BARANG ==========")

	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].id == id {

			ketemu = true

			for {

				fmt.Println("===== PILIH DATA YANG DIUBAH =====")
				fmt.Println("1. Nama")
				fmt.Println("2. Kategori")
				fmt.Println("3. Harga")
				fmt.Println("4. Stok")
				fmt.Println("0. Selesai")

				fmt.Print("Pilih : ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {

					fmt.Print("Nama Baru : ")
					fmt.Scan(&databarang[i].nama)

				} else if pilihan == 2 {

					fmt.Print("Kategori Baru : ")
					fmt.Scan(&databarang[i].kategori)

				} else if pilihan == 3 {

					fmt.Print("Harga Baru : ")
					fmt.Scan(&databarang[i].harga)

				} else if pilihan == 4 {

					fmt.Print("Stok Baru : ")
					fmt.Scan(&databarang[i].stok)

				} else if pilihan == 0 {

					fmt.Println("Edit selesai")
					savebarang()
					break

				} else {

					fmt.Println("Menu tidak tersedia")
				}

				fmt.Println()
			}
		}
	}

	if ketemu == false {

		fmt.Println("Barang tidak ditemukan")
	}
}

func hapusbarang() {

	var id int
	var i int
	var j int
	var ketemu bool

	ketemu = false

	fmt.Println("========== HAPUS BARANG ==========")

	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahbarang; i++ {

		if databarang[i].id == id {

			ketemu = true

			for j = i; j < jumlahbarang-1; j++ {    

				databarang[j] = databarang[j+1]
			}

			jumlahbarang--

			fmt.Println("Barang berhasil dihapus")
			savebarang()
		}
	}

	if ketemu == false {

		fmt.Println("Barang tidak ditemukan")
	}
}

func savebarang() {

	var file *os.File
	var data string
	var i int

	file, _ = os.Create("barang.txt")

	defer file.Close()

	for i = 0; i < jumlahbarang; i++ {

		data =
			strconv.Itoa(databarang[i].id) + "|" +
				databarang[i].nama + "|" +
				databarang[i].kategori + "|" +
				strconv.Itoa(databarang[i].harga) + "|" +
				strconv.Itoa(databarang[i].stok) + "|" +
				strconv.Itoa(databarang[i].terjual) + "\n"

		file.WriteString(data)
	}
}

func loadbarang() {

	var file *os.File
	var scanner *bufio.Scanner
	var line string
	var data []string
	var barang barang

	file, _ = os.Open("barang.txt")

	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {

		line = scanner.Text()

		data = strings.Split(line, "|")

		barang.id, _ =
			strconv.Atoi(data[0])

		barang.nama = data[1]

		barang.kategori = data[2]

		barang.harga, _ =
			strconv.Atoi(data[3])

		barang.stok, _ =
			strconv.Atoi(data[4])

		barang.terjual, _ =
			strconv.Atoi(data[5])

		databarang[jumlahbarang] = barang
		jumlahbarang++
	}
}