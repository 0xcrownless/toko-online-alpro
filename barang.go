package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAXBARANG = 999

type barang struct {
	id       int
	nama     string
	kategori string
	harga    int
	stok     int
	terjual  int
}

var databarang [MAXBARANG]barang
var jumlahbarang int

func tambahbarang() {
	var item barang

	fmt.Println("========== TAMBAH BARANG ==========")
	if jumlahbarang >= MAXBARANG {
		fmt.Println("Data barang sudah penuh")
		return
	}

	item.id = idbarang()
	fmt.Print("Nama Barang : ")
	fmt.Scan(&item.nama)
	fmt.Print("Kategori : ")
	fmt.Scan(&item.kategori)
	fmt.Print("Harga : ")
	fmt.Scan(&item.harga)
	fmt.Print("Stok : ")
	fmt.Scan(&item.stok)

	if item.nama == "" || item.kategori == "" {
		fmt.Println("Nama dan kategori tidak boleh kosong")
		return
	}
	if cariindexbarang(0, item.nama) != -1 {
		fmt.Println("Nama barang sudah digunakan")
		return
	}
	if item.harga <= 0 {
		fmt.Println("Harga harus lebih dari 0")
		return
	}
	if item.stok < 0 {
		fmt.Println("Stok tidak boleh negatif")
		return
	}

	databarang[jumlahbarang] = item
	jumlahbarang++
	savebarang()
	fmt.Println("Barang berhasil ditambahkan")
}

func tampilbarang() {
	urutbarangidasc()
	tampilbarangsaatini()
}

// tampilbarangsaatini menampilkan urutan array saat ini agar hasil sorting tidak diurutkan ulang berdasarkan ID.
func tampilbarangsaatini() {
	fmt.Println("==========================================================================")
	fmt.Println("                           DATA BARANG TOKO")
	fmt.Println("==========================================================================")
	fmt.Println("ID   Nama Barang          Kategori      Harga       Stok        Terjual")
	fmt.Println("--------------------------------------------------------------------------")
	for i := 0; i < jumlahbarang; i++ {
		fmt.Printf("%-4d %-20s %-13s %-11d %-11d %-8d\n",
			databarang[i].id,
			databarang[i].nama,
			databarang[i].kategori,
			databarang[i].harga,
			databarang[i].stok,
			databarang[i].terjual,
		)
	}
	fmt.Println("==========================================================================")
}

func urutbarangidasc() {
	for i := 0; i < jumlahbarang-1; i++ {
		for j := i + 1; j < jumlahbarang; j++ {
			if databarang[j].id < databarang[i].id {
				databarang[i], databarang[j] = databarang[j], databarang[i]
			}
		}
	}
}

func editbarang() {
	var id int

	fmt.Println("========== EDIT BARANG ==========")
	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)

	for i := 0; i < jumlahbarang; i++ {
		if databarang[i].id != id {
			continue
		}
		for {
			var pilihan int
			fmt.Println("===== PILIH DATA YANG DIUBAH =====")
			fmt.Println("1. Nama")
			fmt.Println("2. Kategori")
			fmt.Println("3. Harga")
			fmt.Println("4. Stok")
			fmt.Println("0. Selesai")
			fmt.Print("Pilih : ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				var nama string
				fmt.Print("Nama Baru : ")
				fmt.Scan(&nama)
				index := cariindexbarang(0, nama)
				if nama == "" {
					fmt.Println("Nama tidak boleh kosong")
				} else if index != -1 && index != i {
					fmt.Println("Nama barang sudah digunakan")
				} else if nama != databarang[i].nama && barangadaditransaksi(databarang[i].id, databarang[i].nama) {
					fmt.Println("Nama barang yang sudah tercatat dalam transaksi tidak dapat diubah")
				} else {
					databarang[i].nama = nama
				}
			case 2:
				fmt.Print("Kategori Baru : ")
				fmt.Scan(&databarang[i].kategori)
			case 3:
				var harga int
				fmt.Print("Harga Baru : ")
				fmt.Scan(&harga)
				if harga <= 0 {
					fmt.Println("Harga harus lebih dari 0")
				} else {
					databarang[i].harga = harga
				}
			case 4:
				var stok int
				fmt.Print("Stok Baru : ")
				fmt.Scan(&stok)
				if stok < 0 {
					fmt.Println("Stok tidak boleh negatif")
				} else {
					databarang[i].stok = stok
				}
			case 0:
				savebarang()
				fmt.Println("Edit selesai")
				return
			default:
				fmt.Println("Menu tidak tersedia")
			}
			fmt.Println()
		}
	}
	fmt.Println("Barang tidak ditemukan")
}

func hapusbarang() {
	var id int

	fmt.Println("========== HAPUS BARANG ==========")
	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)

	for i := 0; i < jumlahbarang; i++ {
		if databarang[i].id != id {
			continue
		}
		if barangadaditransaksi(databarang[i].id, databarang[i].nama) {
			fmt.Println("Barang masih tercatat dalam transaksi dan tidak dapat dihapus")
			return
		}
		for j := i; j < jumlahbarang-1; j++ {
			databarang[j] = databarang[j+1]
		}
		jumlahbarang--
		databarang[jumlahbarang] = barang{}
		savebarang()
		fmt.Println("Barang berhasil dihapus")
		return
	}
	fmt.Println("Barang tidak ditemukan")
}

func savebarang() {
	file, err := os.Create("barang.txt")
	if err != nil {
		fmt.Println("Gagal menyimpan data barang")
		return
	}
	defer file.Close()

	for i := 0; i < jumlahbarang; i++ {
		data := strconv.Itoa(databarang[i].id) + "|" +
			databarang[i].nama + "|" +
			databarang[i].kategori + "|" +
			strconv.Itoa(databarang[i].harga) + "|" +
			strconv.Itoa(databarang[i].stok) + "|" +
			strconv.Itoa(databarang[i].terjual) + "\n"
		if _, err := file.WriteString(data); err != nil {
			fmt.Println("Gagal menyimpan data barang")
			return
		}
	}
}

func idbarang() int {
	idterbesar := 0
	for i := 0; i < jumlahbarang; i++ {
		if databarang[i].id > idterbesar {
			idterbesar = databarang[i].id
		}
	}
	return idterbesar + 1
}

func loadbarang() {
	file, err := os.Open("barang.txt")
	if err != nil {
		return
	}
	defer file.Close()

	jumlahbarang = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "|")
		if len(data) < 6 || jumlahbarang >= MAXBARANG {
			continue
		}

		id, errID := strconv.Atoi(data[0])
		harga, errHarga := strconv.Atoi(data[3])
		stok, errStok := strconv.Atoi(data[4])
		terjual, errTerjual := strconv.Atoi(data[5])
		if errID != nil || errHarga != nil || errStok != nil || errTerjual != nil || id <= 0 || harga < 0 || stok < 0 || terjual < 0 {
			continue
		}

		databarang[jumlahbarang] = barang{id: id, nama: data[1], kategori: data[2], harga: harga, stok: stok, terjual: terjual}
		jumlahbarang++
	}
}
