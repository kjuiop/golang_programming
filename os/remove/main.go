package main

import (
	"fmt"
	"os"
)

func main() {
	dirPath := "/tmp/path/test"

	if err := os.RemoveAll(dirPath); err != nil {
		fmt.Printf("Error while removing directory: %v\n", err)
		return
	}
	fmt.Printf("Directory %s and its contents are successfully removed.\n", dirPath)
}
