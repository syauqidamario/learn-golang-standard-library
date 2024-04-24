//Generic enables one to create parameters with changeable types without using the any interface

package main

import (
	"fmt"
	"slices"
)

func main(){
	names := []string{"Mary", "George", "Sheldon", "Melissa", "Georgie"}
	values := []int{100, 99, 98, 97, 96}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Georgie"))
	fmt.Println(slices.Contains(names, "Mary"))
}