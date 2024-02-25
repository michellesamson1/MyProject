package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchFiles(directory, searchString string) {
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".txt") {
			if containsString(path, searchString) {
				fmt.Println("Found in file:", path)
			}
		}

		return nil
	})
}

func containsString(filePath, searchString string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchString) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	return false
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory> <searchString>")
		os.Exit(1)
	}

	directory := os.Args[1]
	searchString := os.Args[2]

	searchFiles(directory, searchString)
}
