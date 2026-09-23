package main

import (
	"fmt"
)

func main() {
	biryani := food{"non-veg" , "spicy"}

	fmt.Println(biryani)
	fmt.Printf("%+v\n" , biryani)

	fmt.Printf("category : %v , taste : %v\n" , biryani.category , biryani.taste)

	
	
}

type food struct {
	category string
	taste string

}
