package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAXTRANSAKSI = 999

type transaksi struct {
	idtransaksi      int
	pembeli          string
	idbarang         int
	namabarang       string
	jumlah           int
	total            int
	status           string
	metodepembayaran string
}

var datatransaksi [MAXTRANSAKSI]transaksi
var jumlahtransaksi int

func idtransaksibaru() int {
	idterbesar := 0
	for i := 0; i < jumlahtransaksi; i++ {
		if datatransaksi[i].idtransaksi > idterbesar {
			idterbesar = datatransaksi[i].idtransaksi
		}
	}
	return idterbesar + 1
}

func pilihmetodepembayaran(pilihan int) (string, bool) {
	switch pilihan {
	case 1:
		return "Dana", true
	case 2:
		return "Ovo", true
	case 3:
		return "Gopay", true
	case 4:
		return "QRIS", true
	case 5:
		return "Bank BCA", true
	case 6:
		return "Bank BRI", true
	case 7:
		return "Bank Mandiri", true
	default:
		return "", false
	}
}

func tampilpilihanpembayaran() {
	fmt.Println("========== PILIH PEMBAYARAN ==========")
	fmt.Println("1. Dana")
	fmt.Println("2. Ovo")
	fmt.Println("3. Gopay")
	fmt.Println("4. QRIS")
	fmt.Println("5. Bank BCA")
	fmt.Println("6. Bank BRI")
	fmt.Println("7. Bank Mandiri")
}

func tambahtransaksi() {
	buattransaksi(namapembeli)
}

func tambahtransaksiadmin() {
	var pembeli string
	fmt.Print("Nama Pembeli : ")
	fmt.Scan(&pembeli)
	if pembeli == "" {
		fmt.Println("Nama pembeli tidak boleh kosong")
		return
	}
	buattransaksi(pembeli)
}

func buattransaksi(pembeli string) {
	var trx transaksi
	var namabarang string
	var pilihanbayar int

	if jumlahtransaksi >= MAXTRANSAKSI {
		fmt.Println("Data transaksi sudah penuh")
		return
	}

	trx.idtransaksi = idtransaksibaru()
	trx.pembeli = pembeli

	fmt.Println("========== TRANSAKSI PEMBELIAN ==========")
	fmt.Print("Nama Barang : ")
	fmt.Scan(&namabarang)
	fmt.Print("Jumlah Beli : ")
	fmt.Scan(&trx.jumlah)

	if trx.jumlah <= 0 {
		fmt.Println("Jumlah beli harus lebih dari 0")
		return
	}

	index := cariindexbarang(0, namabarang)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan")
		return
	}
	if trx.jumlah > databarang[index].stok {
		fmt.Println("Stok tidak cukup")
		return
	}

	tampilpilihanpembayaran()
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihanbayar)
	metode, valid := pilihmetodepembayaran(pilihanbayar)
	if !valid {
		fmt.Println("Metode pembayaran tidak tersedia")
		return
	}

	trx.idbarang = databarang[index].id
	trx.namabarang = databarang[index].nama
	trx.total = trx.jumlah * databarang[index].harga
	trx.status = "pending"
	trx.metodepembayaran = metode

	datatransaksi[jumlahtransaksi] = trx
	jumlahtransaksi++
	savetransaksi()
	fmt.Println("Transaksi berhasil ditambahkan")
}

func approvetransaksi() {
	var idtrx int

	fmt.Println("========== APPROVE TRANSAKSI ==========")
	fmt.Print("Masukkan ID Transaksi : ")
	fmt.Scan(&idtrx)

	for i := 0; i < jumlahtransaksi; i++ {
		if datatransaksi[i].idtransaksi != idtrx {
			continue
		}
		if datatransaksi[i].status == "approved" {
			fmt.Println("Transaksi sudah diapprove")
			return
		}
		if datatransaksi[i].status != "pending" {
			fmt.Println("Status transaksi tidak valid")
			return
		}
		if datatransaksi[i].jumlah <= 0 {
			datatransaksi[i].jumlah = hitungjumlahbarangtransaksi(datatransaksi[i].namabarang)
		}
		if !prosesstoktransaksi(datatransaksi[i], false) {
			return
		}
		if !prosesstoktransaksi(datatransaksi[i], true) {
			return
		}

		datatransaksi[i].status = "approved"
		saldotoko += datatransaksi[i].total
		savetransaksi()
		savebarang()
		fmt.Println("Transaksi berhasil diapprove")
		return
	}
	fmt.Println("Transaksi tidak ditemukan")
}

func tampiltransaksi() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                   DATA TRANSAKSI                                          ║")
	fmt.Println("╠════╦════════════╦════════════════════╦════════════╦════════════╦══════════════════════════╣")
	fmt.Println("║ ID ║ Pembeli    ║ Barang             ║ Total      ║ Status     ║ Pembayaran               ║")
	fmt.Println("╠════╬════════════╬════════════════════╬════════════╬════════════╬══════════════════════════╣")
	for i := 0; i < jumlahtransaksi; i++ {
		fmt.Printf("║ %-2d ║ %-10s ║ %-18s ║ %-10d ║ %-10s ║ %-24s ║\n",
			datatransaksi[i].idtransaksi,
			datatransaksi[i].pembeli,
			datatransaksi[i].namabarang,
			datatransaksi[i].total,
			datatransaksi[i].status,
			datatransaksi[i].metodepembayaran,
		)
	}
	fmt.Println("╚════╩════════════╩════════════════════╩════════════╩══════════════════════════╝")
}

func savetransaksi() {
	if err := simpandaftartransaksi("transaksi.txt", daftartransaksiaktif()); err != nil {
		fmt.Println("Gagal menyimpan transaksi")
	}
}

func daftartransaksiaktif() []transaksi {
	daftar := make([]transaksi, jumlahtransaksi)
	for i := 0; i < jumlahtransaksi; i++ {
		daftar[i] = datatransaksi[i]
	}
	return daftar
}

func formatbaristransaksi(trx transaksi) string {
	return strconv.Itoa(trx.idtransaksi) + "|" +
		trx.pembeli + "|" +
		strconv.Itoa(trx.idbarang) + "|" +
		trx.namabarang + "|" +
		strconv.Itoa(trx.jumlah) + "|" +
		strconv.Itoa(trx.total) + "|" +
		trx.status + "|" +
		trx.metodepembayaran + "\n"
}

func simpandaftartransaksi(namafile string, daftar []transaksi) error {
	file, err := os.Create(namafile)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < len(daftar); i++ {
		if _, err := file.WriteString(formatbaristransaksi(daftar[i])); err != nil {
			return err
		}
	}
	return nil
}

func loadtransaksi() {
	file, err := os.Open("transaksi.txt")
	if err != nil {
		return
	}
	defer file.Close()

	jumlahtransaksi = 0
	saldotoko = loadpendapatanarsipbulanan()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "|")
		if len(data) < 8 || jumlahtransaksi >= MAXTRANSAKSI {
			continue
		}

		id, errID := strconv.Atoi(data[0])
		idBarang, errBarang := strconv.Atoi(data[2])
		jumlah, errJumlah := strconv.Atoi(data[4])
		total, errTotal := strconv.Atoi(data[5])
		if errID != nil || errBarang != nil || errJumlah != nil || errTotal != nil || id <= 0 || jumlah <= 0 || total < 0 {
			continue
		}
		if data[6] != "pending" && data[6] != "approved" {
			continue
		}

		trx := transaksi{
			idtransaksi:      id,
			pembeli:          data[1],
			idbarang:         idBarang,
			namabarang:       data[3],
			jumlah:           jumlah,
			total:            total,
			status:           data[6],
			metodepembayaran: data[7],
		}
		datatransaksi[jumlahtransaksi] = trx
		jumlahtransaksi++
		if trx.status == "approved" {
			saldotoko += trx.total
		}
	}
}

func hapustransaksi() {
	var id int

	fmt.Println("========== HAPUS TRANSAKSI ==========")
	fmt.Print("Masukkan ID transaksi : ")
	fmt.Scan(&id)

	for i := 0; i < jumlahtransaksi; i++ {
		if datatransaksi[i].idtransaksi != id {
			continue
		}
		if datatransaksi[i].status == "approved" {
			fmt.Println("Transaksi approved tidak dapat dihapus")
			return
		}
		for j := i; j < jumlahtransaksi-1; j++ {
			datatransaksi[j] = datatransaksi[j+1]
		}
		jumlahtransaksi--
		datatransaksi[jumlahtransaksi] = transaksi{}
		savetransaksi()
		fmt.Println("Transaksi berhasil dihapus")
		return
	}
	fmt.Println("Transaksi tidak ditemukan")
}
