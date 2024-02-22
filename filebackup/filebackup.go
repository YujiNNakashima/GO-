package filebackup

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func HashIt() {
	// Get the filename from command line arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	h := sha1.New()

	h.Write(data)

	hashResult := h.Sum(nil)

	hashString := fmt.Sprintf("%x", hashResult)

	fmt.Printf("SHA1 of \"%s\" is %s\n", filename, hashString)
}
