package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func parse(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
        return nil, fmt.Errorf("failed to open file: %v", err)
    }
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: %v", err)
	}

	if len(records) == 0 {
		fmt.Println("CSV file is empty")
	}

	return records, nil
}
