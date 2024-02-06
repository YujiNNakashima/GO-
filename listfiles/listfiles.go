package listfiles

import (
	"fmt"
	"os"
)

func ListFiles() {
	fmt.Println(os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <srcDir>")
		os.Exit(1)
	}

	srcDir := os.Args[1]
	files, err := os.ReadDir(srcDir)
	if err != nil {
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
