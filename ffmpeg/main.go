package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// FFmpeg를 사용하여 비디오 파일을 변환하는 예제
	// 여기서 "input.mp4"를 "output.mp4"로 변환합니다
	cmd := exec.Command("ffmpeg", "-i", "./input.mp4", "./output.mp4")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("ffmpeg cmd fail, err : %s\n", err.Error())
	}
	fmt.Println(string(output))
}

// docker build -t go-ffmpeg-app .
// docker run --rm go-ffmpeg-app
