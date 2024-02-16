package countthelines

import (
	"bufio"
	"os"
	"path/filepath"
)

func CountLinesInFiles(globPattern string) (int, error) {
	// Glob for files matching the provided pattern
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
	}

	return totalLines, nil
}
