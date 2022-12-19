package pkg

import (
	"encoding/csv"
	"log"
	"os"
)


func WriteToFile(s [][]string) {
	f, err := os.Create("output.csv")
	defer f.Close()
	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()
	for _, records := range s {
		w.Write(records)
	}

}