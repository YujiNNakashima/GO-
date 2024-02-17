// main.go
package main

import (
	// "fmt"
	"fmt"
	"os"
	"softwaredesign/countthelines"
)

func main() {
	// listfiles.ListFiles()
	// asyncgo.AsyncGoo(2112)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <globPattern>")
		return
	}

	_, err := countthelines.CountLinesInFiles(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
