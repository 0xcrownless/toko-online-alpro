package main 
import "fmt"

func main (){
	var role string 
	var pilihan int 

	initadmin()
	role = login()

	if role == "superadmin" {
		for pilihan != 0 {
			fmt.Println("===== SUPER ADMIN AKSES =====")
			fmt.Println("1. tambah admin")
			fmt.Println("2. lihat admin")
			fmt.Println("3. Masuk sistem toko ")
			fmt.Println("0. keluar")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				tambahadmin()
			}else if pilihan == 2 {
				tampiladmin()

			}else if pilihan == 3 {
				menutoko()
			}else if pilihan == 0 {
				fmt.Println("progrsm selesai")
			}else {
				fmt.Println("menu tidak tersedia ")
			}
			fmt.Println()
		}
	} else if role == "admin" {
		menutoko()
	}
}

func menutoko() {
	var pilihan int 

	for pilihan != 0 {
		fmt.Println("========== TOko ONLINE ==========")
		fmt.Println("1. tambah")
		fmt.Println("2. tampilkan barang ")
		fmt.Println("3. edit barang ")
		fmt.Println("4. hapus barang")
		fmt.Println("0. keluar")
		fmt.Print("pilih menu :")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahbarang()
		} else if pilihan == 2 {
			tampilbarang()
		} else if pilihan == 3 {
			editbarang()
		} else if pilihan == 4 {
			hapusbarang()
		} else if pilihan == 0 {
			fmt.Println("kembali")
		} else {
			fmt.Println("menu tidak ada")
		}
		fmt.Println()
	}
}