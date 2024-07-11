package main

import (
	"fmt"
	"unsafe"
)

type Sockaddr_in struct {
	Sin_len    uint8
	Sin_family uint8
	Sin_port   uint16
	Sin_addr   In_addr
	Sin_zero   [8]int8
}

type In_addr struct {
	S_addr uint8
}

func main() {
	fmt.Println(unsafe.Sizeof(Sockaddr_in{}))
}
