package main

import (
	"fmt"
	"softwaredesign/filebackup"
	"time"
)

func main() {
	// listfiles.ListFiles()
	// asyncgo.AsyncGoo(2112)
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: go run main.go <globPattern>")
	// 	return
	// }

	// _, err := countthelines.CountLinesInFiles(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	rootDir := "./countthelines"
	// pairs, err := filebackup.HashIt(rootDir)
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }

	// fmt.Println("File paths e hashes:")
	// for _, pair := range pairs {
	// 	fmt.Printf("%s: %s\n", pair[0], pair[1])
	// }

	// newFiles, err := filebackup.FindNew(rootDir, pairs)
	// if err != nil {
	// 	panic(err)
	// }

	// // Print new files
	// for hash, path := range newFiles {
	// 	println(hash, ":", path)
	// }

	// src := "/path/to/source"
	dst := "./countthelines/bkp"
	timestamp := time.Now().Unix()

	if err := filebackup.BackItUp(rootDir, dst, timestamp); err != nil {
		fmt.Println("Backup failed:", err)
		return
	}

	// fmt.Println("Backup completed successfully")
}
