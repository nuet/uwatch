package excel

import (
	"github.com/tealeg/xlsx"
	"strconv"
)

//导入
func Import(filename string, cellmaps []string, args ...int) (data []map[string]interface{}, err error) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return
	}

	mapslen := len(cellmaps)
	startColumn := 1 //起始列，默认第1列 即A列
	endColumn := 26  //结束列,默认第26列 即Z列
	startRow := 1    //起始行，默认第1行
	argsnum := len(args)
	if argsnum > 0 {
		if args[0] > 0 {
			startColumn = args[0]
		}

		if argsnum > 1 {
			endColumn = args[1]
			if argsnum > 2 && args[2] > 0 {
				startRow = args[2]
			}
		}
	}

	var key string
	for _, sheet := range xlFile.Sheets {
		rowsNum := len(sheet.Rows)

		for r := startRow - 1; r < rowsNum; r++ {
			cellsNum := len(sheet.Rows[r].Cells)
			if cellsNum == 0 {
				continue
			}

			row := make(map[string]interface{})
			emptyNum := 0
			for c := startColumn - 1; c < endColumn; c++ {
				if mapslen > c && cellmaps[c] != "" {
					key = cellmaps[c]
				} else {
					key = strconv.Itoa(c)
				}
				if c >= cellsNum {
					emptyNum++
					row[key] = ""
					continue
				}
				if sheet.Rows[r].Cells[c].Value == "" {
					emptyNum++
				}
				row[key] = sheet.Rows[r].Cells[c].Value
			}
			if endColumn-startColumn+1 > emptyNum {
				data = append(data, row)
			}
		}
	}
	return
}
