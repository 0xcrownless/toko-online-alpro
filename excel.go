package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func exportexcel() {

	var i int

	var omzet int
	var totaltransaksi int
	var totalbarang int

	file := excelize.NewFile()

	// DASHBOARD

	dashboard := "Dashboard"

	file.NewSheet(dashboard)

	for i = 0; i < jumlahtransaksi; i++ {

		if datatransaksi[i].status == "approved" {

			omzet += datatransaksi[i].total
			totaltransaksi++
			totalbarang += datatransaksi[i].jumlah
		}
	}

	file.SetCellValue(dashboard, "A1", "DASHBOARD TOKO")

	file.SetCellValue(dashboard, "A3", "Saldo Toko")
	file.SetCellValue(dashboard, "B3", saldotoko)

	file.SetCellValue(dashboard, "A4", "Omzet Toko")
	file.SetCellValue(dashboard, "B4", omzet)

	file.SetCellValue(dashboard, "A5", "Total Transaksi")
	file.SetCellValue(dashboard, "B5", totaltransaksi)

	file.SetCellValue(dashboard, "A6", "Barang Terjual")
	file.SetCellValue(dashboard, "B6", totalbarang)

	//BARANG

	barangsheet := "Barang"

	file.NewSheet(barangsheet)

	file.SetCellValue(barangsheet, "A1", "ID")
	file.SetCellValue(barangsheet, "B1", "Nama Barang")
	file.SetCellValue(barangsheet, "C1", "Harga")
	file.SetCellValue(barangsheet, "D1", "Stok")

	for i = 0; i < jumlahbarang; i++ {

		file.SetCellValue(
			barangsheet,
			"A"+fmt.Sprint(i+2),
			databarang[i].id,
		)

		file.SetCellValue(
			barangsheet,
			"B"+fmt.Sprint(i+2),
			databarang[i].nama,
		)

		file.SetCellValue(
			barangsheet,
			"C"+fmt.Sprint(i+2),
			databarang[i].harga,
		)

		file.SetCellValue(
			barangsheet,
			"D"+fmt.Sprint(i+2),
			databarang[i].stok,
		)
	}

	//TRANSAKSI

	transaksi := "Transaksi"

	file.NewSheet(transaksi)

	file.SetCellValue(transaksi, "A1", "ID")
	file.SetCellValue(transaksi, "B1", "Pembeli")
	file.SetCellValue(transaksi, "C1", "Barang")
	file.SetCellValue(transaksi, "D1", "Total")
	file.SetCellValue(transaksi, "E1", "Pembayaran")
	file.SetCellValue(transaksi, "F1", "Status")

	for i = 0; i < jumlahtransaksi; i++ {

		file.SetCellValue(
			transaksi,
			"A"+fmt.Sprint(i+2),
			datatransaksi[i].idtransaksi,
		)

		file.SetCellValue(
			transaksi,
			"B"+fmt.Sprint(i+2),
			datatransaksi[i].pembeli,
		)

		file.SetCellValue(
			transaksi,
			"C"+fmt.Sprint(i+2),
			datatransaksi[i].namabarang,
		)

		file.SetCellValue(
			transaksi,
			"D"+fmt.Sprint(i+2),
			datatransaksi[i].total,
		)

		file.SetCellValue(
			transaksi,
			"E"+fmt.Sprint(i+2),
			datatransaksi[i].metodepembayaran,
		)

		file.SetCellValue(
			transaksi,
			"F"+fmt.Sprint(i+2),
			datatransaksi[i].status,
		)
	}

	// CHART

	file.AddChart(
		barangsheet,
		"F2",
		&excelize.Chart{
			Type: excelize.Col,
			Series: []excelize.ChartSeries{
				{
					Name:       "Penjualan",
					Categories: "Barang!$B$2:$B$20",
					Values:     "Barang!$D$2:$D$20",
				},
			},
			Title: []excelize.RichTextRun{
				{
					Text: "Grafik Stok Barang",
				},
			},
		},
	)

	file.DeleteSheet("Sheet1")

	err := file.SaveAs("laporan.xlsx")

	if err != nil {

		fmt.Println("Export gagal")
		return
	}

	fmt.Println("Export Excel berhasil")
}