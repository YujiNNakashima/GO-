package filebackup

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func HashIt() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := args[1]

	// read in chunks
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	h := sha1.New()

	// copia direto ao objeto de hash, sem carregar tudo na memória
	_, err = io.Copy(h, file)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}

	hashResult := h.Sum(nil)

	hashString := hex.EncodeToString(hashResult)

	fmt.Printf("SHA1 of \"%s\" is %s\n", filename, hashString)
}
