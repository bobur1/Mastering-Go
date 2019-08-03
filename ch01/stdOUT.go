package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {

	} else {
		myString = arguments[1]
	}
	fmt.Println("hera is argument:",arguments,"<--")
	fmt.Println("argument length:",len(arguments))
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}