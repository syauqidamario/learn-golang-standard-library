package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Syauqi")
	data.PushBack("Damario")
	data.PushBack("Djohan")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}