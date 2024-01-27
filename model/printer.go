package model

import "fmt"

func PrintEcho(data []byte) []byte {
	fmt.Println(string(data))
	return data
}
