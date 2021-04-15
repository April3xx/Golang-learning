package main

import (
	"encoding/hex"
	"fmt"
)

func test() {
	key := "6368616e676520746869732070617373"
	fmt.Print("the key length is :")
	fmt.Println(len(key))
	fmt.Print("converted to byte")
	fmt.Println(hex.DecodeString(key))
	temp, _ := hex.DecodeString(key)
	fmt.Println(string(temp))
}
