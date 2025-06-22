package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check() string{
	if len(os.Args) < 2 {
		fmt.Println("Usage: gcsv <filePath>")
		os.Exit(1)
		return ""
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	return filepath.Base(filePath)
}
