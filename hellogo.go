package main

import (
	"fmt"
)

func main() {
	today := 'r'
	switch {
	case today == 'r':
		fmt.Println("r")
	case today == 'w':
		fmt.Println("w")
	default:
		fmt.Println("error")
	}
}
