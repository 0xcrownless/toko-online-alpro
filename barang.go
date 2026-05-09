package main 
import "fmt"

const MAXBARANG = 999
type barang struct {
	id int
	nama string
	kategori string
	harga int
	stok int 
	terjual int 
}

var databrang [MAXBARANG]barang
var jumlahbarang int 

func tambahbarang() {
	var barang barang 

	fmt.Println("======== SILAHKAN TAMBAH BARANG ===========")
	fmt.Print("id barang  ; ")
	fmt.Scam(&barang.id)
	fmt.Print("nama barang : ")
	fmt.Scan(&barang.nama)
	fmt.Print("kategori ; ")
	fmt.Scan(&barang.kategori)
	fmt.Print("harga : ")
	fmt.Scan(&barang.harga)
	fmt.Print("stok : ")
	fmt.Scan(&barang.stok)

	barang.terjual = 0
	
	databarang[jumlahbarang] = barang 
	jumlahbarang++

	fmt.Println("barang berhasil ditambahkan")
}