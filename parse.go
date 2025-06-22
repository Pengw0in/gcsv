package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func parse() [][]string {

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read csv file: %v", err)
	}

	if len(records) == 0 {
		fmt.Println("[INF]CSV file is empty")
	}

	return records
}
