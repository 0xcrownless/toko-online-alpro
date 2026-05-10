package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func exportexcel() {

	var file *excelize.File

	file = excelize.NewFile()

	fmt.Println("Excel berhasil dibuat")
}