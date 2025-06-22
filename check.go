package main

import (
	"path/filepath"
	"fmt"
	"log"
	"os"
)

func check() string {
	if len(os.Args) < 2 {
		fmt.Println("[ERR]Usage: gcsv <fileName>")
	}

	
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	return filepath.Base(filePath)
}
