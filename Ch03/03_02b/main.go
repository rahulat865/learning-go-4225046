package main

import (
	"fmt"
)

func main() {
	anInt := 10

	var p* int = &anInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p : " , *p)
	}

	value1 := 10.23
	pointer1 := &value1
	*pointer1 = *pointer1 + 10
	*pointer1 = *pointer1 + 10

	fmt.Println("value of pointer : " ,*pointer1)
	fmt.Println(" original value : " ,value1)
	
}
