package main

import (
	"fmt"
	"sort"
)

func main() {
	var mkc = make(map[string]string)
	mkc["AWS"] = "amazon web services"
	mkc["micrsoft"] = "azure"
	mkc["google"] = "cloud"

	fmt.Println(mkc)

	ec := mkc["AWS"]
	fmt.Println(ec)

	mkc["oracle"] = "no cloud"
	
	
	// delete the value 
	delete(mkc , "google")
	fmt.Println(mkc)

	// displaying the values of map through loop

	fmt.Println("displaying the values of map through loop")

	for k , v := range mkc{
		fmt.Printf("%v: %v\n" , k ,v)
	}

	// to show key values of map in alphabetical order

	keys := make([]string , len(mkc))
	i := 0 
	for k := range mkc{
		keys[i] = k
	i++
	}

	sort.Strings(keys)
	fmt.Println("\nSorted order ")
	
	for i := range keys{
		fmt.Println(mkc[keys[i]])
	}

 }
