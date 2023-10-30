package main

import (
	"fmt"
	likedlist "go_leti_hw1/likedList"
)

func main() {
	x := likedlist.New(10)
	fmt.Println(x.Size())
	x.Print()
	x.Add(100)
	x.Print()
	x.DeleteFrom(5)
	x.Print()
	x.UpdateVal(10, 2)
	x.Print()
	fmt.Println(x.GetVal(2))

	var values []int
	values = append(values, 1, 2, 3, 4, 5)
	fmt.Println(values)
	linkFromValue := likedlist.LikedListFromSlice(values)
	linkFromValue.Print()

}
