package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func TestRecords() {
	//Read CSV and generate a string [][]
	logger := log.Default()
	logger.Print("Populate database running....")

	records := GetRecords("./mock/company.csv", ",")
	logger.Print("🌿", records)
}

func GetRecords(path, separator string) [][]string {
	records := [][]string{}
	if len(path) == 0 {
		return records
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return records
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(separator)[0]
	for {
		columnValues, err := csvReader.Read()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		records = append(records, columnValues)
	}
	return records
}
