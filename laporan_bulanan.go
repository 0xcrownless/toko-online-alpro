package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

const filependapatanarsipbulanan = "pendapatan_arsip_bulanan.txt"

type baranglaporanbulanan struct {
	nama     string
	kategori string
	harga    int
	jumlah   int
	total    int
}

func tutuplaporanbulanan() {
	waktu := time.Now()
	approved, pending := pisahtransaksibulanan()

	if len(approved) == 0 && len(pending) == 0 {
		fmt.Println("Belum ada transaksi untuk ditutup")
		return
	}

	fileExcel, err := exportlaporanbulananexcel(waktu, approved, pending)
	if err != nil {
		fmt.Println("Gagal membuat laporan Excel bulanan:", err)
		fmt.Println("Reset transaksi dibatalkan")
		return
	}

	fileTXT, err := buatarsipbulananTXT(waktu, approved, pending)
	if err != nil {
		fmt.Println("Gagal membuat arsip TXT:", err)
		fmt.Println("Reset transaksi dibatalkan")
		return
	}

	pendapatanApproved, _ := ringkasantransaksibulanan(approved)
	pendapatanArsipLama := loadpendapatanarsipbulanan()
	if err := simpanpendapatanarsipbulanan(pendapatanArsipLama + pendapatanApproved); err != nil {
		fmt.Println("Gagal menyimpan pendapatan arsip:", err)
		fmt.Println("Reset transaksi dibatalkan")
		return
	}

	if err := simpandaftartransaksi("transaksi.txt", pending); err != nil {
		_ = simpanpendapatanarsipbulanan(pendapatanArsipLama)
		fmt.Println("Gagal mereset transaksi:", err)
		fmt.Println("Reset transaksi dibatalkan")
		return
	}

	datatransaksi = [MAXTRANSAKSI]transaksi{}
	for i := 0; i < len(pending); i++ {
		datatransaksi[i] = pending[i]
	}
	jumlahtransaksi = len(pending)

	fmt.Println("Tutup laporan bulanan berhasil")
	fmt.Println("Arsip TXT  :", fileTXT)
	fmt.Println("Excel      :", fileExcel)
	fmt.Println("Transaksi approved sudah diarsipkan")
	fmt.Println("Transaksi pending tetap tersimpan")
}

func pisahtransaksibulanan() ([]transaksi, []transaksi) {
	var approved []transaksi
	var pending []transaksi

	for i := 0; i < jumlahtransaksi; i++ {
		if datatransaksi[i].status == "approved" {
			approved = append(approved, datatransaksi[i])
		} else if datatransaksi[i].status == "pending" {
			pending = append(pending, datatransaksi[i])
		}
	}
	return approved, pending
}

func folderlaporanbulanan() string {
	return "laporan_bulanan"
}

func namafilelaporanbulanan(waktu time.Time, ekstensi string) string {
	nama := fmt.Sprintf("laporan_%s.%s", waktu.Format("2006_01"), ekstensi)
	return filepath.Join(folderlaporanbulanan(), nama)
}

func exportlaporanbulananexcel(waktu time.Time, approved []transaksi, pending []transaksi) (string, error) {
	if err := os.MkdirAll(folderlaporanbulanan(), 0755); err != nil {
		return "", err
	}

	namafile := namafilelaporanbulanan(waktu, "xlsx")
	if _, err := os.Stat(namafile); err == nil {
		var jawab string
		fmt.Print("File laporan Excel bulan ini sudah ada. Timpa file lama? (y/n) ")
		fmt.Scan(&jawab)
		if strings.ToLower(jawab) != "y" {
			return "", fmt.Errorf("pembuatan laporan Excel dibatalkan oleh pengguna")
		}
	} else if !os.IsNotExist(err) {
		return "", err
	}

	if err := buatExcelLaporanBulanan(namafile, waktu, approved, pending); err != nil {
		return "", err
	}
	return namafile, nil
}

func buatExcelLaporanBulanan(namafile string, waktu time.Time, approved []transaksi, pending []transaksi) error {
	file := excelize.NewFile()
	defer file.Close()

	ringkasan := "Ringkasan"
	if err := file.SetSheetName("Sheet1", ringkasan); err != nil {
		return err
	}

	barangTerjual := hitungbarangterjualbulanan(approved)
	topBarang := salinbaranglaporanbulanan(barangTerjual)
	urutbaranglaporanbulanan(topBarang)

	isisheetringkasan(file, ringkasan, waktu, approved, pending)
	isisheettransaksibulanan(file, "Transaksi Approved", approved)
	isisheettransaksibulanan(file, "Transaksi Pending", pending)
	isisheetbarangbulanan(file, "Barang Terjual", barangTerjual)
	isisheetbarangbulanan(file, "Top Barang", topBarang)

	if err := file.SaveAs(namafile); err != nil {
		return err
	}
	return nil
}

func isisheetringkasan(file *excelize.File, sheet string, waktu time.Time, approved []transaksi, pending []transaksi) {
	totalPendapatan, totalItem := ringkasantransaksibulanan(approved)

	file.SetCellValue(sheet, "A1", "LAPORAN BULANAN TOKO")
	file.SetCellValue(sheet, "A3", "Bulan dan Tahun")
	file.SetCellValue(sheet, "B3", fmt.Sprintf("%02d/%04d", int(waktu.Month()), waktu.Year()))
	file.SetCellValue(sheet, "A4", "Jumlah Transaksi Approved")
	file.SetCellValue(sheet, "B4", len(approved))
	file.SetCellValue(sheet, "A5", "Jumlah Transaksi Pending")
	file.SetCellValue(sheet, "B5", len(pending))
	file.SetCellValue(sheet, "A6", "Total Penjualan Approved")
	file.SetCellValue(sheet, "B6", totalPendapatan)
	file.SetCellValue(sheet, "A7", "Total Item Terjual")
	file.SetCellValue(sheet, "B7", totalItem)
	file.SetCellValue(sheet, "A8", "Waktu Laporan Ditutup")
	file.SetCellValue(sheet, "B8", waktu.Format("2006-01-02 15:04:05"))
	file.SetColWidth(sheet, "A", "B", 28)
}

func isisheettransaksibulanan(file *excelize.File, sheet string, daftar []transaksi) {
	file.NewSheet(sheet)
	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Pembeli")
	file.SetCellValue(sheet, "C1", "ID Barang")
	file.SetCellValue(sheet, "D1", "Barang")
	file.SetCellValue(sheet, "E1", "Jumlah")
	file.SetCellValue(sheet, "F1", "Total")
	file.SetCellValue(sheet, "G1", "Pembayaran")
	file.SetCellValue(sheet, "H1", "Status")

	for i := 0; i < len(daftar); i++ {
		baris := fmt.Sprint(i + 2)
		file.SetCellValue(sheet, "A"+baris, daftar[i].idtransaksi)
		file.SetCellValue(sheet, "B"+baris, daftar[i].pembeli)
		file.SetCellValue(sheet, "C"+baris, daftar[i].idbarang)
		file.SetCellValue(sheet, "D"+baris, daftar[i].namabarang)
		file.SetCellValue(sheet, "E"+baris, daftar[i].jumlah)
		file.SetCellValue(sheet, "F"+baris, daftar[i].total)
		file.SetCellValue(sheet, "G"+baris, daftar[i].metodepembayaran)
		file.SetCellValue(sheet, "H"+baris, daftar[i].status)
	}
	file.SetColWidth(sheet, "A", "H", 18)
}

func isisheetbarangbulanan(file *excelize.File, sheet string, daftar []baranglaporanbulanan) {
	file.NewSheet(sheet)
	file.SetCellValue(sheet, "A1", "Nama Barang")
	file.SetCellValue(sheet, "B1", "Kategori")
	file.SetCellValue(sheet, "C1", "Harga")
	file.SetCellValue(sheet, "D1", "Jumlah Terjual")
	file.SetCellValue(sheet, "E1", "Total Pendapatan")

	for i := 0; i < len(daftar); i++ {
		baris := fmt.Sprint(i + 2)
		file.SetCellValue(sheet, "A"+baris, daftar[i].nama)
		file.SetCellValue(sheet, "B"+baris, daftar[i].kategori)
		file.SetCellValue(sheet, "C"+baris, daftar[i].harga)
		file.SetCellValue(sheet, "D"+baris, daftar[i].jumlah)
		file.SetCellValue(sheet, "E"+baris, daftar[i].total)
	}
	file.SetColWidth(sheet, "A", "E", 20)
}

func ringkasantransaksibulanan(daftar []transaksi) (int, int) {
	var totalPendapatan int
	var totalItem int

	for i := 0; i < len(daftar); i++ {
		totalPendapatan += daftar[i].total
		if daftar[i].jumlah > 0 {
			totalItem += daftar[i].jumlah
		} else {
			totalItem += hitungjumlahbarangtransaksi(daftar[i].namabarang)
		}
	}
	return totalPendapatan, totalItem
}

func hitungbarangterjualbulanan(daftar []transaksi) []baranglaporanbulanan {
	var hasil []baranglaporanbulanan
	indexBarang := make(map[string]int)

	for i := 0; i < len(daftar); i++ {
		trx := daftar[i]
		if trx.idbarang != 0 {
			jumlah := trx.jumlah
			if jumlah <= 0 {
				jumlah = hitungjumlahbarangtransaksi(trx.namabarang)
			}
			tambahbaranglaporanbulanan(&hasil, indexBarang, trx.namabarang, jumlah, trx.total)
			continue
		}

		for _, item := range strings.Split(trx.namabarang, ",") {
			nama, jumlah, berhasil := bacaitemkeranjang(item)
			if !berhasil {
				continue
			}
			harga := hargabaranglaporan(nama)
			tambahbaranglaporanbulanan(&hasil, indexBarang, nama, jumlah, harga*jumlah)
		}
	}
	return hasil
}

func tambahbaranglaporanbulanan(hasil *[]baranglaporanbulanan, indexBarang map[string]int, nama string, jumlah int, total int) {
	if strings.TrimSpace(nama) == "" || jumlah <= 0 {
		return
	}

	kategori := "-"
	harga := 0
	index := cariindexbarang(0, nama)
	if index != -1 {
		kategori = databarang[index].kategori
		harga = databarang[index].harga
	}
	if total <= 0 {
		total = harga * jumlah
	}

	posisi, ada := indexBarang[nama]
	if ada {
		(*hasil)[posisi].jumlah += jumlah
		(*hasil)[posisi].total += total
		return
	}

	indexBarang[nama] = len(*hasil)
	*hasil = append(*hasil, baranglaporanbulanan{
		nama:     nama,
		kategori: kategori,
		harga:    harga,
		jumlah:   jumlah,
		total:    total,
	})
}

func hargabaranglaporan(nama string) int {
	index := cariindexbarang(0, nama)
	if index == -1 {
		return 0
	}
	return databarang[index].harga
}

func salinbaranglaporanbulanan(daftar []baranglaporanbulanan) []baranglaporanbulanan {
	salinan := make([]baranglaporanbulanan, len(daftar))
	for i := 0; i < len(daftar); i++ {
		salinan[i] = daftar[i]
	}
	return salinan
}

func urutbaranglaporanbulanan(daftar []baranglaporanbulanan) {
	for i := 0; i < len(daftar)-1; i++ {
		for j := i + 1; j < len(daftar); j++ {
			if daftar[j].jumlah > daftar[i].jumlah ||
				(daftar[j].jumlah == daftar[i].jumlah && daftar[j].total > daftar[i].total) {
				daftar[i], daftar[j] = daftar[j], daftar[i]
			}
		}
	}
}

func buatarsipbulananTXT(waktu time.Time, approved []transaksi, pending []transaksi) (string, error) {
	if err := os.MkdirAll(folderlaporanbulanan(), 0755); err != nil {
		return "", err
	}

	namafile := namafilelaporanbulanan(waktu, "txt")
	totalPendapatan, totalItem := ringkasantransaksibulanan(approved)

	var isi strings.Builder
	isi.WriteString("LAPORAN BULANAN TOKO\n")
	isi.WriteString("=====================\n")
	isi.WriteString("Bulan dan Tahun: " + fmt.Sprintf("%02d/%04d", int(waktu.Month()), waktu.Year()) + "\n")
	isi.WriteString("Waktu Ditutup: " + waktu.Format("2006-01-02 15:04:05") + "\n")
	isi.WriteString("Jumlah Transaksi Approved: " + strconv.Itoa(len(approved)) + "\n")
	isi.WriteString("Jumlah Transaksi Pending: " + strconv.Itoa(len(pending)) + "\n")
	isi.WriteString("Total Penjualan Approved: " + strconv.Itoa(totalPendapatan) + "\n")
	isi.WriteString("Total Item Terjual: " + strconv.Itoa(totalItem) + "\n\n")

	isi.WriteString("TRANSAKSI APPROVED\n")
	isi.WriteString("ID|Pembeli|ID Barang|Nama Barang|Jumlah|Total|Status|Metode Pembayaran\n")
	for i := 0; i < len(approved); i++ {
		isi.WriteString(formatbaristransaksi(approved[i]))
	}

	isi.WriteString("\nTRANSAKSI PENDING\n")
	isi.WriteString("ID|Pembeli|ID Barang|Nama Barang|Jumlah|Total|Status|Metode Pembayaran\n")
	for i := 0; i < len(pending); i++ {
		isi.WriteString(formatbaristransaksi(pending[i]))
	}

	if err := os.WriteFile(namafile, []byte(isi.String()), 0644); err != nil {
		return "", err
	}
	return namafile, nil
}

func loadpendapatanarsipbulanan() int {
	data, err := os.ReadFile(filependapatanarsipbulanan)
	if err != nil {
		return 0
	}

	pendapatan, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil || pendapatan < 0 {
		return 0
	}
	return pendapatan
}

func simpanpendapatanarsipbulanan(pendapatan int) error {
	return os.WriteFile(filependapatanarsipbulanan, []byte(strconv.Itoa(pendapatan)+"\n"), 0644)
}
