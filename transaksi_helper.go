package main

import (
	"fmt"
	"strconv"
	"strings"
)

// prosesstoktransaksi memeriksa semua item terlebih dahulu dan baru mengubah stok bila seluruh item valid.
// Barang yang sama di keranjang dijumlahkan agar tidak dapat lolos melebihi stok.
func prosesstoktransaksi(trx transaksi, ubah bool) bool {
	var jumlahperbarang [MAXBARANG]int

	if trx.idbarang != 0 {
		if trx.jumlah <= 0 {
			fmt.Println("Jumlah beli harus lebih dari 0")
			return false
		}
		index := cariindexbarang(trx.idbarang, trx.namabarang)
		if index == -1 {
			fmt.Println("Barang transaksi tidak ditemukan")
			return false
		}
		jumlahperbarang[index] = trx.jumlah
	} else {
		if strings.TrimSpace(trx.namabarang) == "" {
			fmt.Println("Data barang transaksi tidak sesuai")
			return false
		}
		for _, item := range strings.Split(trx.namabarang, ",") {
			nama, jumlah, berhasil := bacaitemkeranjang(item)
			if !berhasil {
				fmt.Println("Data barang transaksi tidak sesuai")
				return false
			}
			index := cariindexbarang(0, nama)
			if index == -1 {
				fmt.Println("Barang", nama, "tidak ditemukan")
				return false
			}
			jumlahperbarang[index] += jumlah
		}
	}

	for i := 0; i < jumlahbarang; i++ {
		if jumlahperbarang[i] > databarang[i].stok {
			fmt.Println("Stok", databarang[i].nama, "tidak cukup")
			return false
		}
	}

	if ubah {
		for i := 0; i < jumlahbarang; i++ {
			if jumlahperbarang[i] > 0 {
				databarang[i].stok -= jumlahperbarang[i]
				databarang[i].terjual += jumlahperbarang[i]
			}
		}
	}
	return true
}

func cariindexbarang(idbarang int, namabarang string) int {
	for i := 0; i < jumlahbarang; i++ {
		if idbarang != 0 && databarang[i].id == idbarang {
			return i
		}
	}
	for i := 0; i < jumlahbarang; i++ {
		if databarang[i].nama == namabarang {
			return i
		}
	}
	return -1
}

func bacaitemkeranjang(item string) (string, int, bool) {
	item = strings.TrimSpace(item)
	posisi := strings.LastIndex(item, " x")
	if posisi == -1 {
		return "", 0, false
	}

	nama := strings.TrimSpace(item[:posisi])
	jumlah, err := strconv.Atoi(strings.TrimSpace(item[posisi+2:]))
	if err != nil || nama == "" || jumlah <= 0 {
		return "", 0, false
	}
	return nama, jumlah, true
}

func hitungjumlahbarangtransaksi(namabarang string) int {
	total := 0
	for _, item := range strings.Split(namabarang, ",") {
		_, jumlah, berhasil := bacaitemkeranjang(item)
		if berhasil {
			total += jumlah
		}
	}
	return total
}

func barangadaditransaksi(id int, nama string) bool {
	for i := 0; i < jumlahtransaksi; i++ {
		trx := datatransaksi[i]
		if trx.idbarang != 0 && trx.idbarang == id {
			return true
		}
		for _, item := range strings.Split(trx.namabarang, ",") {
			namaitem, _, berhasil := bacaitemkeranjang(item)
			if berhasil && namaitem == nama {
				return true
			}
		}
	}
	return false
}
