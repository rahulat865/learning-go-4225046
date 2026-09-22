package main

import "fmt"

func main() {
	i1 , i2 , i3 := 6 , 69 , 12
	intSum := i1 + i2 + i3
 	fmt.Println("Integer Sum :", intSum)

	f1 , f2 , f3 := 6.5 , 69.5 , 12.5
	floatSum := f1 + f2 + f3
  fmt.Println("The float sum is : " , floatSum)


	// sum := i1 + f1  this will not work because we are trying to add an integer and a float

	total := float64(i1) + f3
	fmt.Println("The result is : ", total)
}
