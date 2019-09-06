package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "abcde"
	fmt.Println(strings.Index(line, "f"))
}
