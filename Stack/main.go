package main

import (
	"fmt"
	"strconv"

	linkedList "github.com/RioChndr/go-learn/linked-list"
)

func main(){
	list := linkedList.List{}


	for i := 0; i < 10; i++ {
		list.Push(&linkedList.Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-data-" + strconv.Itoa(i),
		})
	}

	fmt.Printf("Size %d on init\n", list.Size)

	fmt.Printf("Head id %v\n", list.Head.Id)

	list.Pop()

	fmt.Printf("Size %d after pop\n", list.Size)
	fmt.Printf("Head id %v after pop\n", list.Head.Id)
}