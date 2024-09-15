package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ApplicationCodeCounter(printEveryFile bool) {
	root := "." // starting directory
	totalLines := 0

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only Go files
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			lines, err := countLines(path)
			if err != nil {
				return err
			}
			if printEveryFile {
				fmt.Printf("%s: %d lines\n", path, lines)
			}
			totalLines += lines
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return
	}

	fmt.Printf("\nTotal lines of code: %d\n", totalLines)
}

// countLines counts the non-empty and non-comment lines in a file
func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	inBlockComment := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue // skip empty lines
		}

		// Handle block comments
		if strings.HasPrefix(line, "/*") {
			inBlockComment = true
		}
		if inBlockComment {
			if strings.HasSuffix(line, "*/") {
				inBlockComment = false
			}
			continue
		}

		// Skip single line comments
		if strings.HasPrefix(line, "//") {
			continue
		}

		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
