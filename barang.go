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

func tampilbarang() {
	var i int 
	fmt.Println("=========== data barang ========")
	
	if jumlah barang == 0 {
		fmt.Println("data barang kosong ")

	}else {
		for i 0; i < jumlahbarang; i++ {

			fmt.Println("")
			fmt.Println("id : ", databarang[i].id)
			fmt.Println("nama : ", databarang[i].nama)
			fmt.Println("kategori : ", databarang[i].kategori)
			fmt.Println("harga : ", databarang[i].harga)
			fmt.Println("stok : ", databarang[i].stok)
			fmt.Println("terjual: ", databarang[i].terjual)
			fmt.Println("")
		}
	}
}

func editbarang() {
	var id, i int 
	var ketemu bool 

	ketemu = false 

	fmt.Println("========== edit barang ==========")
	fmt.Print("masukan id barang : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahbarang; i++ {
		if databarang[i].id == id {
			ketemu = true 

			fmt.Print("nama baru : ")
			fmt.Scan(&databarang[i].nama)
			fmt.Print("kategori baru : ")
			fmt.Scan(&databarang[i].kategori)
			fmt.Print("harga baru : ")
			fmt.Scan(&databarang[i].harga)
			fmt.Print("stok baru : ")
			fmt.Scan(&databarang[i].stok)
			fmt.Println("data berhasil diubah")
		}
	}

	if ketemu == false {
		fmt.Println("barang tidak ditemukan")
	}
}

func hapusbarang(){
	var id, i, j int 
	var ketemu bool 

	ketemu = false 

	fmt.Println("======= hapus barang =======")
	fmt.Print("masukan id barang : ")
	fmt.Scan(&id)

	for i = 0; i < jumlahbarang; i++ {
		if databarang[i].id == id {
			ketemu = true

			for j = i; j < jumlahbarang-1; j++ {
				databarang[j] = databarang[j+1]
			}
			jumlahbarang--
			fmt.Println("barang barhasil dihapus")
		}
	}

	if ketemu == false {
		fmt.Println("barang tidak ditemukan")
	}

}