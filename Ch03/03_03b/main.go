package main

import (
	"fmt"
)

func main() {
	var colors [5]string 

	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	colors[3] = "Yellow"
	colors[4] = "Purple"

	fmt.Println("number of colors : " , len(colors))
	fmt.Println("List of Colors are : ", colors)

	var numerics = [5]int{32 , 56 , 69 , 81 , 93}
	fmt.Println("All the values of numerics :" ,numerics)
	fmt.Println("All numerics are : " , len(numerics))

}
