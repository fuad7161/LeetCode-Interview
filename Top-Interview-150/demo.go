package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "     This    is  a   string   with   irregular  spacing   "
	fmt.Println(line)

	line = strings.TrimSpace(line)
	word := strings.Fields(line)
	fmt.Println(len(word))
	result := strings.Join(word, " ")
	fmt.Println(result)
}
