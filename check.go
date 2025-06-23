package main

import (
	"fmt"
	"os"
)

func check() (string, error){
	if len(os.Args) < 2 {
		return "", fmt.Errorf("usage: gcsv <filePath>")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		return  "", fmt.Errorf("failed to open file: %v", err)
	}
	file.Close()
	return filePath, nil
}
