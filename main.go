package main

import (
	"pdf_reader/pkg"

	"log"
	"path/filepath"
)

func main() {
	var fullRecords = [][]string{}
	fullRecords = append(fullRecords, []string{"PO Number", "Total Net Amount"})

	pdf, err := filepath.Glob("./PutAllPDFFilesHere/*.pdf")
	if err != nil {
		log.Fatal(err)
	}
	for _, eachPDF := range pdf {
		content, err := pkg.ReadPdf(eachPDF) // Read local pdf file
		if err != nil {
			panic(err)
		}

		fullRecords = append(fullRecords, content)
	}

	pkg.WriteToFile(fullRecords)

}
