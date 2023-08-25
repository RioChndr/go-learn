package main

import (
	"fmt"
	"strconv"

	linkedList "github.com/RioChndr/go-learn/linked-list"
)

func main(){
	list := linkedList.List{}


	for i := 0; i < 10; i++ {
		list.Add(&linkedList.Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-data-" + strconv.Itoa(i),
		})
	}

	fmt.Printf("Size %d on init\n", list.Size)

	fmt.Printf("Last id %v\n", list.Tail.Id)

	list.Remove(list.Tail.Id)

	fmt.Printf("Size %d after remove last\n", list.Size)
	fmt.Printf("Last id %v after remove last\n", list.Tail.Id)
}