package LinkedList

import (
	"strconv"
	"testing"
)

func TestAddMultipleItem(t *testing.T){
	var link List = List{}

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
	if(link.Head.Child.Id != "Test2"){
		t.Fatalf("Expected child id is 'Test2'")
	}

	var _,item = link.Find("Test2")
	if(item.Data != "Test2-data"){
		t.Fatalf("Expected childd ata is Test2-data")
	}
}

func TestAddMoreData(t *testing.T){
	var list = List{}
	
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
	var list = List{}
	
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
	var list = List{}

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
	if(list.Head.Child.Id != "new-after-0"){
		t.Fatalf("After head should be new-after-0")
	}

	list.AddAfter(list.Tail.Id, &Item{
		Id: "last-after-99",
		Data: "last-after-99",
	})
	if(list.Tail.Id != "last-after-99"){
		t.Fatalf("After head should be last-after-99")
	}

	if(list.Size != 10 + 2){
		t.Fatalf("size should be 102")
	}
}

func TestRemove(t *testing.T){
	var list = List{}

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
	list.Remove(list.Tail.Id)
	if list.Size != 7 {
		t.Fatalf("Size should be 7")
	}
	list.Remove("unknown-id")
	if list.Size != 7 {
		t.Fatalf("Size should be 7 if no id found")
	}

	list.Remove(list.Head.Child.Id)
	if list.Size != 6 {
		t.Fatalf("Size should be 6 after delete inner id")
	}
}

func TestRemoveLast(t *testing.T){
	list := List{}


	for i := 0; i < 10; i++ {
		list.Add(&Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-data-" + strconv.Itoa(i),
		})
	}

	list.Remove(list.Tail.Id)

	if(list.Size != 9){
		t.Fatal("List size should 9 after remove tail")
	}

	if(list.Tail.Id == "new-9"){
		t.Fatal("Tile id should not new-9")
	}

	if(list.Tail.Id != "new-8"){
		t.Fatal("Tile id should new-8")
	}
}

func TestPush(t *testing.T){
	list := List{}

	if list.Size != 0{
		t.Fatalf("list should have size 0 on first init")
	}

	list.Push(&Item{
		Id: "First push",
		Data: "First data push",
	})
	if list.Head.Id != "First push"{
		t.Fatalf("list head should be First push")
	}
	if list.Size != 1{
		t.Fatalf("list should have size 1")
	}

	
	list.Push(&Item{
		Id: "Second push",
		Data: "Second push data",
	})
	if list.Size != 2{
		t.Fatalf("list should have size 2")
	}
	if(list.Head.Id != "Second push"){
		t.Fatal("Head id should second push")
	}
}

func TestPop(t *testing.T){
	var list = List{}
	
	for i := 0; i < 100; i++ {
		list.Add(&Item{
			Id: "new-" + strconv.Itoa(i),
			Data: "new-" + strconv.Itoa(i),
		})
	}

	list.Pop()

	if(list.Head.Id != "new-1"){
		t.Fatal("Head after pop should be new-1")
	}
}