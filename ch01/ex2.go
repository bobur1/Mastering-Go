package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	//Assume we will write only float numbers in command line
	arguments := os.Args
	sum := 0.0
	amount := 0.0
	for index, element := range arguments {
		if(index == 0){
			continue
		}

		amount ++
		float_element, _ := strconv.ParseFloat(element, 64)
		sum += float_element
	}

	fmt.Println("Avarage = ",sum/amount)

}
