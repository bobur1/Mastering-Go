package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	arguments := os.Args
	sum :=0
	for index, element := range arguments {
		if(index == 0){
			continue
		}

		int_element, _ := strconv.Atoi(element)
		sum += int_element
	}
	fmt.Println("Sum = ",sum)

}
