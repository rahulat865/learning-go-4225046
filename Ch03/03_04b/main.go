package main

import (
	"fmt"
	"sort"
	
)

func main() {
	// This is an array
	var colors = make([]string ,  0 , 4)
	colors = append(colors , "Red" , "Green" , "magenta")
	fmt.Println(colors)
	colors = append(colors, "Purple" , "Blue")
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)

	colors = remove(colors , 1)
	fmt.Println(colors)
}

func remove(slice []string , i int) []string {
	return append(slice[:i] , slice[i+1:]...)
}
