package countthelines

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func worker(filePath string, resultChan chan<- fileResult, errorChan chan<- error) {
	file, err := os.Open(filePath)
	if err != nil {
		errorChan <- err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		errorChan <- err
		return
	}

	resultChan <- fileResult{filePath, lineCount}
}

type fileResult struct {
	filePath string
	lines    int
}

func CountLinesInFiles(globPattern string) (int, error) {
	matchedFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return 0, err
	}

	resultChan := make(chan fileResult)
	errorChan := make(chan error)

	for _, filePath := range matchedFiles {
		go worker(filePath, resultChan, errorChan)
	}

	totalLines := 0
	for i := 0; i < len(matchedFiles); i++ {
		select {
		case result := <-resultChan:
			fmt.Printf("Total number of lines from %s: %d\n", result.filePath, result.lines)
			totalLines += result.lines
		case err := <-errorChan:
			return 0, err
		}
	}

	return totalLines, nil
}
