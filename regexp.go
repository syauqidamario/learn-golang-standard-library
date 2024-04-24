package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("syauqi"))
	fmt.Println(regex.MatchString("syauki"))
	fmt.Println(regex.MatchString("syouQi"))
	fmt.Println(regex.FindAllString("syauqi syaki syauQi", 10))
}