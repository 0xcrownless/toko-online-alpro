package main

import "fmt"

var saldotoko int

func lihatsaldo() {

	fmt.Println("========== SALDO TOKO ==========")
	fmt.Println("Saldo toko : Rp", saldotoko)
}

func tariksaldo() {

	var jumlah int

	fmt.Println("========== TARIK SALDO ==========")

	fmt.Print("Jumlah tarik : Rp ")
	fmt.Scan(&jumlah)

	if jumlah > saldotoko {

		fmt.Println("Saldo tidak cukup")
		return
	}

	saldotoko -= jumlah

	fmt.Println("Saldo berhasil ditarik")
	fmt.Println("Sisa saldo : Rp", saldotoko)
}