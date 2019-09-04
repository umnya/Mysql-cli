package main

import (
    "bufio"
    "fmt"
	"os"
//	"log"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}