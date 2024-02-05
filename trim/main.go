package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "2000000 / 2000000"
	if strings.Contains(str, "/") {
		strBitrate := strings.Split(str, "/")
		trimBitrate := strings.TrimSpace(strBitrate[0])
		fmt.Printf("trimmed Bitrate : %s\n", trimBitrate)
		if bitrate, err := strconv.ParseInt(trimBitrate, 10, 64); err == nil {
			fmt.Printf("[success] audio bitrate Bitrate : %d", bitrate)
		}
	}
}
