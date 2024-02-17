package countthelines

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func CountLinesInFiles(globPattern string) (int, error) {

	matchedFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return 0, err
	}

	totalLines := 0

	for _, filePath := range matchedFiles {
		file, err := os.Open(filePath)
		if err != nil {
			return 0, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			totalLines++
		}

		if err := scanner.Err(); err != nil {
			return 0, err
		}
		fmt.Printf("Total number of lines from %s: %d\n", filePath, totalLines)
	}

	return totalLines, nil
}
