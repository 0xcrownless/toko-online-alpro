package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var saldotoko int
var totalpenarikan int

func loadpenarikan() {
	data, err := os.ReadFile("penarikan.txt")
	if err != nil {
		return
	}
	jumlah, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil || jumlah < 0 {
		return
	}
	totalpenarikan = jumlah
	saldotoko -= totalpenarikan
}

func savepenarikan() {
	if err := os.WriteFile("penarikan.txt", []byte(strconv.Itoa(totalpenarikan)+"\n"), 0644); err != nil {
		fmt.Println("Gagal menyimpan data penarikan")
	}
}

func lihatsaldo() {
	fmt.Println("========== SALDO TOKO ==========")
	fmt.Println("Saldo toko : Rp", saldotoko)
}

func tariksaldo() {
	var jumlah int

	fmt.Println("========== TARIK SALDO ==========")
	fmt.Print("Jumlah tarik : Rp ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah tarik harus lebih dari 0")
		return
	}
	if jumlah > saldotoko {
		fmt.Println("Saldo tidak cukup")
		return
	}

	saldotoko -= jumlah
	totalpenarikan += jumlah
	savepenarikan()
	fmt.Println("Saldo berhasil ditarik")
	fmt.Println("Sisa saldo : Rp", saldotoko)
}
