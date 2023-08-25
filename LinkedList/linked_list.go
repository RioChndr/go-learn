package LinkedList

// Linked list

type Item struct {
	Child *Item
	Id string
	Data string
}

type List struct {
	Head *Item
	Tail *Item
	Size int
}

func (link *List) Add(newItem *Item){
	if(link.Tail == nil){
		link.Head = newItem
		link.Tail = newItem
	} else {
		link.Tail.Child = newItem
		link.Tail = newItem
	}
	link.Size = link.Size + 1
}

func (link *List) Find(id string) (before *Item, current *Item){
	var indexBefore *Item
	var indexCurr *Item = link.Head;
	var selectedItem *Item
	for indexCurr != nil {
		if(indexCurr.Id == id){
			selectedItem = indexCurr
			break
		}
		indexBefore = indexCurr
		indexCurr = indexCurr.Child
	}
	return indexBefore, selectedItem
}

func (link *List) AddAfter(id string, newItem *Item){
	var _,selectedCurr = link.Find(id)
	if selectedCurr == nil {
		return;
	}
	if selectedCurr == link.Tail{
		link.Tail = newItem
	} else {
		newItem.Child = selectedCurr.Child
	}
	selectedCurr.Child = newItem
	link.Size = link.Size + 1
}

func (link *List) Remove(id string){
	var selectedBefore,selectedCurr = link.Find(id)
	if selectedCurr == nil {
		return;
	}
	link.Size = link.Size - 1
	if selectedCurr == link.Head {
		link.Head = selectedCurr.Child
		return
	}
	if selectedCurr == link.Tail {
		selectedBefore.Child = nil
		link.Tail = selectedBefore
		return
	}
	selectedBefore.Child = selectedCurr.Child
}

func (l *List) Push(i *Item){
	if l.Head != nil{
		i.Child = l.Head
	}
	
	l.Head = i
	l.Size = l.Size + 1
}

func (l *List) Pop() (*Item){
	old := l.Head
	l.Head = l.Head.Child	
	l.Size = l.Size - 1
	return old;
}
