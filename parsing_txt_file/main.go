package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "/tmp/result_all_count.txt"

	// 파일 열기
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("파일을 열 수 없습니다: %v", err)
		return
	}
	defer file.Close()

	uniqueIds := make(map[string]LogStruct)
	// 파일 내용을 라인별로 읽기
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		if !strings.Contains(line, "N5_01H1QRH5EES1V6QECX82VSYPH8") {
			//fmt.Println(line)
			continue
		}

		var logStruct LogStruct
		err := json.Unmarshal([]byte(line), &logStruct)
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return
		}

		userId := parsingUserId(line)
		_, exists := uniqueIds[userId]
		if !exists {
			uniqueIds[userId] = logStruct
			fmt.Printf("user_id: %s, time : %s, count : %d\n", userId, logStruct.Time, len(uniqueIds))
		}
	}

	// 에러 처리
	if err := scanner.Err(); err != nil {
		fmt.Printf("파일을 읽는 도중 에러가 발생했습니다: %v", err)
		return
	}

	fmt.Println("Set size:", len(uniqueIds))
}

func parsingUserId(line string) string {
	index := strings.Index(line, "user=")
	if index >= 0 {
		line := line[index+len("user="):]
		lastIndex := strings.Index(line, "stream")
		userId := line[:lastIndex-7]
		return userId
	}

	return ""
}

type LogStruct struct {
	Log    string `json:"log"`
	Stream string `json:"stream"`
	Time   string `json:"time"`
}
