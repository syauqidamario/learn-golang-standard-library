package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Syauqi Damario", "Syauqi"))
	fmt.Println(strings.Split("Syauqi Damario", ""))
	fmt.Println(strings.ToLower("Syauqi Damario"))
	fmt.Println(strings.ToUpper("Syauqi Damario"))
	fmt.Println(strings.Trim("     Syauqi Damario   ", ""))
	fmt.Println(strings.ReplaceAll("Syauqi Syauqi Syauqi Syauqi", "Syauqi", "Damario"))
}