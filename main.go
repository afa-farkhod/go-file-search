package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <directory-or-file> <keyword>")
		os.Exit(1)
	}

	target := os.Args[1]
	keyword := os.Args[2]

	info, err := os.Stat(target)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if info.IsDir() {
		err := filepath.Walk(target, func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error accessing %s: %v\n", path, err)
				return nil
			}
			if !fi.IsDir() {
				searchFile(path, keyword)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Walk error: %v\n", err)
		}
	} else {
		searchFile(target, keyword)
	}
}

func searchFile(path, keyword string) {
	file, err := os.Open(path)
	if err != nil {
		return // silently skip unreadable files
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Printf("[MATCH] %s (line %d): %s\n", path, lineNumber, line)
		}
		lineNumber++
	}
}