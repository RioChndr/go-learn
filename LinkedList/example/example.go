package main

import (
	"fmt"
	"math/rand"
	"strconv"

	linkedList "github.com/RioChndr/go-learn/linked-list"
)

func main(){
	var list = linkedList.List{}
	
	for i := 0; i < 100; i++ {
		list.Add(&linkedList.Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-" + strconv.Itoa(i),
		})
	}

	// Insert after
	for i := 0; i < 100; i++ {
		list.AddAfter("new-" + strconv.Itoa(i), &linkedList.Item{
			Id: "new-after-" + strconv.Itoa(i),
			Data: "new-after-" + strconv.Itoa(i),
		})
	}

	// Insert after random
	for i := 0; i < 100; i++ {
		var randomInt = rand.Intn(100)
		list.AddAfter("new-after-" + strconv.Itoa(randomInt), &linkedList.Item{
			Id: "new-after-random-" + strconv.Itoa(i),
			Data: "new-after-random-" + strconv.Itoa(i),
		})
	}

	var headItem = list.Head
	for headItem.Child != nil {
		fmt.Printf(headItem.Data + "\n")
		headItem = headItem.Child
	}
}