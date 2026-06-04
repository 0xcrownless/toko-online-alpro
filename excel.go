package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func exportexcel() {
	var totalpendapatan int
	var totaltransaksi int
	var totalbarang int

	for i := 0; i < jumlahtransaksi; i++ {
		if datatransaksi[i].status == "approved" {
			totalpendapatan += datatransaksi[i].total
			totaltransaksi++
			totalbarang += datatransaksi[i].jumlah
		}
	}

	file := excelize.NewFile()
	defer file.Close()

	// DASHBOARD
	dashboard := "Dashboard"
	file.NewSheet(dashboard)
	file.SetCellValue(dashboard, "A1", "DASHBOARD TOKO")
	file.SetCellValue(dashboard, "A3", "Saldo Toko")
	file.SetCellValue(dashboard, "B3", saldotoko)
	file.SetCellValue(dashboard, "A4", "Total Penjualan")
	file.SetCellValue(dashboard, "B4", totalpendapatan)
	file.SetCellValue(dashboard, "A5", "Total Penarikan")
	file.SetCellValue(dashboard, "B5", totalpenarikan)
	file.SetCellValue(dashboard, "A6", "Total Transaksi")
	file.SetCellValue(dashboard, "B6", totaltransaksi)
	file.SetCellValue(dashboard, "A7", "Barang Terjual")
	file.SetCellValue(dashboard, "B7", totalbarang)

	// BARANG
	barangsheet := "Barang"
	file.NewSheet(barangsheet)
	file.SetCellValue(barangsheet, "A1", "ID")
	file.SetCellValue(barangsheet, "B1", "Nama Barang")
	file.SetCellValue(barangsheet, "C1", "Kategori")
	file.SetCellValue(barangsheet, "D1", "Harga")
	file.SetCellValue(barangsheet, "E1", "Stok")
	file.SetCellValue(barangsheet, "F1", "Terjual")
	for i := 0; i < jumlahbarang; i++ {
		baris := fmt.Sprint(i + 2)
		file.SetCellValue(barangsheet, "A"+baris, databarang[i].id)
		file.SetCellValue(barangsheet, "B"+baris, databarang[i].nama)
		file.SetCellValue(barangsheet, "C"+baris, databarang[i].kategori)
		file.SetCellValue(barangsheet, "D"+baris, databarang[i].harga)
		file.SetCellValue(barangsheet, "E"+baris, databarang[i].stok)
		file.SetCellValue(barangsheet, "F"+baris, databarang[i].terjual)
	}

	// TRANSAKSI
	transaksisheet := "Transaksi"
	file.NewSheet(transaksisheet)
	file.SetCellValue(transaksisheet, "A1", "ID")
	file.SetCellValue(transaksisheet, "B1", "Pembeli")
	file.SetCellValue(transaksisheet, "C1", "Barang")
	file.SetCellValue(transaksisheet, "D1", "Jumlah")
	file.SetCellValue(transaksisheet, "E1", "Total")
	file.SetCellValue(transaksisheet, "F1", "Pembayaran")
	file.SetCellValue(transaksisheet, "G1", "Status")
	for i := 0; i < jumlahtransaksi; i++ {
		baris := fmt.Sprint(i + 2)
		file.SetCellValue(transaksisheet, "A"+baris, datatransaksi[i].idtransaksi)
		file.SetCellValue(transaksisheet, "B"+baris, datatransaksi[i].pembeli)
		file.SetCellValue(transaksisheet, "C"+baris, datatransaksi[i].namabarang)
		file.SetCellValue(transaksisheet, "D"+baris, datatransaksi[i].jumlah)
		file.SetCellValue(transaksisheet, "E"+baris, datatransaksi[i].total)
		file.SetCellValue(transaksisheet, "F"+baris, datatransaksi[i].metodepembayaran)
		file.SetCellValue(transaksisheet, "G"+baris, datatransaksi[i].status)
	}

	// Grafik memakai kolom Terjual, bukan kolom Stok.
	if jumlahbarang > 0 {
		barisAkhir := jumlahbarang + 1
		err := file.AddChart(barangsheet, "H2", &excelize.Chart{
			Type: excelize.Col,
			Series: []excelize.ChartSeries{{
				Name:       "Barang!$F$1",
				Categories: fmt.Sprintf("Barang!$B$2:$B$%d", barisAkhir),
				Values:     fmt.Sprintf("Barang!$F$2:$F$%d", barisAkhir),
			}},
			Title: []excelize.RichTextRun{{Text: "Grafik Barang Terjual"}},
		})
		if err != nil {
			fmt.Println("Gagal membuat grafik Excel")
			return
		}
	}

	file.DeleteSheet("Sheet1")
	if err := file.SaveAs("laporan.xlsx"); err != nil {
		fmt.Println("Export gagal")
		return
	}
	fmt.Println("Export Excel berhasil")
}
