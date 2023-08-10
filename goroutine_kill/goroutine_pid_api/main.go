package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	memFile, err := os.OpenFile(fmt.Sprintf("/proc/%d/mem", 30984), os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening memory file:", err)
		return
	}
	defer memFile.Close()
	a := uintptr(0xc0000ae008)
	fmt.Println("a : ", a)
	buf := make([]byte, 4)
	_, err = memFile.ReadAt(buf, int64(a))
	if err != nil {
		fmt.Println("error reading memoty :", err)
		return
	}
	value := *(*uint32)(unsafe.Pointer(&buf[0]))
	fmt.Println("value :", value)
}
