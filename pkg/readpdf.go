package pkg

import (
	"fmt"

	"github.com/ledongthuc/pdf"
)

type AppConfig struct {
	OrderNum int
	TotalSum int
}

var App AppConfig

func ReadPdf(path string) ([]string, error) {
	var str []string

	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return nil, err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()

		for _, row := range rows {
			// println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				if word.S == "Purchase Order No." {
					App.OrderNum = 1
					continue
				}
				if App.OrderNum == 1 {
					fmt.Println(word.S)
					str = append(str, word.S)
					App.OrderNum = 0
				}
				if word.S == "Total Net Amount" {
					// fmt.Println(word.S)
					App.TotalSum = 1
					continue
				}
				if App.TotalSum == 1 {
					fmt.Println(word.S)
					str = append(str, word.S)
					App.TotalSum = 0
				}

			}

		}

	}

	return str, nil
}
