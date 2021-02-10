package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	enter()
}
func enter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese palabra a traducir")
	text, _ := reader.ReadString('\n')
	fmt.Println(stringToBin(text))
}
func stringToBin(s string) (binString string) {
	for i, c := range s {
		if i != 0 {
			binString += " "
		}
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return binString
}
