package main

import (
	"github.com/tealeg/xlsx"
)

func main() {
	// read
	filePath := "read.xlsx"
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	sheet := file.Sheets[0]
	for _, row := range sheet.Rows {
		deviceId, _ := row.Cells[1].String()
		accountId, _ := row.Cells[2].String()
		// do something
	}

	// write
	filePath2 := "write.xlsx"
	file2, err := xlsx.OpenFile(filePath2)
	if err != nil {
		panic(err)
	}
	sheet2 := file2.Sheets[0]
	for _, row := range sheet2.Rows {
		cell := row.AddCell()
		cell.Value = ""
	}
	err = file2.Save(filePath2)
	if err != nil {
		panic(err)
	}
}
