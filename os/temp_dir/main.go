package main

import (
	"fmt"
	"log"
	"os"
)

const tempDirNamePattern = "ai-working_*"

func main() {
	dir, err := MakeTempWorkSpace()
	if err != nil {
		log.Println("err : ", err)
	}

	fmt.Println(dir)
}

func MakeTempWorkSpace() (string, error) {
	dir, err := os.MkdirTemp("", tempDirNamePattern)
	if err != nil {
		return "", fmt.Errorf("create temp dir: %w", err)
	}
	return dir, nil
}
