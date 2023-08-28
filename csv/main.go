package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("/tmp/example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// UTF-8 BOM 추가
	bom := []byte{0xEF, 0xBB, 0xBF}
	_, err = file.Write(bom)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"이름", "나이", "직업"}
	writer.Write(header)

	data := []string{"Alice", "30", "개발자"}
	writer.Write(data)

	data = []string{"Bob", "28", "디자이너"}
	writer.Write(data)

	writer.Flush()
}
