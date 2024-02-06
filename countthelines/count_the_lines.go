package countthelines

import (
	"bufio"
	"fmt"
	"os"
)

func CountTheLines() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filePath>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number of lines in %s: %d\n", filePath, lineCount)

}
