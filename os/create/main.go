package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("hi.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fileInfo, _ := f.Stat()
	fmt.Printf("fileInfo : %v\n", fileInfo)
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer
	fmt.Printf("buffer : %v\n", buffer)

	f.Close()
}
