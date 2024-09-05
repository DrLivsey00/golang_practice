package main

import (
	"fmt"
)

func main() {
	res, ok := Answer("What is 5 plus -2?")
	if ok {
		fmt.Println("Answer:", res)
	} else {
		fmt.Println("Error")
	}

}
