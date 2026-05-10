package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func exportexcel() {

	var file *excelize.File

	file = excelize.NewFile()

	file.SaveAs("laporan.xlsx")

	fmt.Println("Excel berhasil dibuat")
}