package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func exportexcel() {

	var file *excelize.File
	var sheet string
	var i int

	file = excelize.NewFile()

	sheet = "Laporan"

	file.SetSheetName("Sheet1", sheet)

	file.SetCellValue(sheet, "A1", "Nama Barang")
	file.SetCellValue(sheet, "B1", "Jumlah Terjual")

	for i = 0; i < jumlahbarang; i++ {

		file.SetCellValue(
			sheet,
			"A"+fmt.Sprint(i+2),
			databarang[i].nama,
		)

		file.SetCellValue(
			sheet,
			"B"+fmt.Sprint(i+2),
			databarang[i].terjual,
		)
	}

	file.AddChart(sheet, "D2", &excelize.Chart{
		Type: "col",
		Series: []excelize.ChartSeries{
			{
				Name:       "Penjualan",
				Categories: "Laporan!$A$2:$A$20",
				Values:     "Laporan!$B$2:$B$20",
			},
		},
		Title: []excelize.RichTextRun{
			{
				Text: "Grafik Penjualan Barang",
			},
		},
	})

	file.SaveAs("laporan.xlsx")

	fmt.Println("Excel dan chart berhasil dibuat")
}