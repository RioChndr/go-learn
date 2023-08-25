package linkedList

import (
	"strconv"
	"testing"
)

func TestAddMultipleItem(t *testing.T){
	var link LinkedList = LinkedList{}

	link.Add(&Item{
		Id: "Test1",
		Data: "Test1",
	})

	if(link.Size != 1){
		t.Fatalf("Expected size to become 1")
	}
	if(link.Head.Id != "Test1"){
		t.Fatalf("Expected head id is 'Test1'")
	}
	
	link.Add(&Item{
		Id: "Test2",
		Data: "Test2-data",
	})

	if(link.Size != 2){
		t.Fatalf("Expected size to become 2")
	}
	if(link.Head.Next.Id != "Test2"){
		t.Fatalf("Expected child id is 'Test2'")
	}

	var _,item = link.Find("Test2")
	if(item.Data != "Test2-data"){
		t.Fatalf("Expected childd ata is Test2-data")
	}
}

func TestAddMoreData(t *testing.T){
	var list = LinkedList{}
	
	for i := 0; i < 100; i++ {
		list.Add(&Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-" + strconv.Itoa(i),
		})
	}

	if list.Size != 100 {
		t.Fatalf("Expected size should 100")
	}
}

func TestFindData(t *testing.T){
	var list = LinkedList{}
	
	for i := 0; i < 10; i++ {
		list.Add(&Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-data-" + strconv.Itoa(i),
		})
	}

	var beforeFirstData,fistData = list.Find("new-0")
	if(fistData.Data != "new-data-0"){
		t.Fatalf("expected first data to be new-data-0")
	}
	if(beforeFirstData != nil){
		t.Fatalf("beforeFirstData should be nil")
	}
}

func TestAddAfter(t *testing.T){
	var list = LinkedList{}

	list.AddAfter("id-unknown", &Item{
		Id: "Should-now-inserted",
	})
	if(list.Head != nil){
		t.Fatalf("Head should be nil")
	}

	for i := 0; i < 10; i++ {
		list.Add(&Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-data-" + strconv.Itoa(i),
		})
	}

	list.AddAfter(list.Head.Id, &Item{
		Id: "new-after-0",
		Data: "new-after-0",
	})
	if(list.Head.Next.Id != "new-after-0"){
		t.Fatalf("After head should be new-after-0")
	}

	list.AddAfter(list.Last.Id, &Item{
		Id: "last-after-99",
		Data: "last-after-99",
	})
	if(list.Last.Id != "last-after-99"){
		t.Fatalf("After head should be last-after-99")
	}

	if(list.Size != 10 + 2){
		t.Fatalf("size should be 102")
	}
}

func TestRemove(t *testing.T){
	var list = LinkedList{}

	list.Remove("no-id-yet")
	if list.Size != 0{
		t.Fatalf("size should be nil")
	}

	for i := 0; i < 10; i++ {
		list.Add(&Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-data-" + strconv.Itoa(i),
		})
	}

	list.Remove(list.Head.Id)
	if list.Size != 9 {
		t.Fatalf("Size should be 9")
	}
	list.Remove(list.Head.Id)
	if list.Size != 8 {
		t.Fatalf("Size should be 8")
	}
	list.Remove(list.Last.Id)
	if list.Size != 7 {
		t.Fatalf("Size should be 7")
	}
	list.Remove("unknown-id")
	if list.Size != 7 {
		t.Fatalf("Size should be 7 if no id found")
	}

	list.Remove(list.Head.Next.Id)
	if list.Size != 6 {
		t.Fatalf("Size should be 6 after delete inner id")
	}
}