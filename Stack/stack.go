package Stack

// Linked list

type Item struct {
	Next *Item
	Id string
	Data string
}

type List struct {
	Head *Item
	Last *Item
	Size int
}

func (link *List) Add(newItem *Item){
	if(link.Last == nil){
		link.Head = newItem
		link.Last = newItem
	} else {
		link.Last.Next = newItem
		link.Last = newItem
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
		indexCurr = indexCurr.Next
	}
	return indexBefore, selectedItem
}

func (link *List) AddAfter(id string, newItem *Item){
	var _,selectedCurr = link.Find(id)
	if selectedCurr == nil {
		return;
	}
	if selectedCurr == link.Last{
		link.Last = newItem
	} else {
		newItem.Next = selectedCurr.Next
	}
	selectedCurr.Next = newItem
	link.Size = link.Size + 1
}

func (link *List) Remove(id string){
	var selectedBefore,selectedCurr = link.Find(id)
	if selectedCurr == nil {
		return;
	}
	link.Size = link.Size - 1
	if selectedCurr == link.Head {
		link.Head = selectedCurr.Next
		return
	}
	if selectedCurr == link.Last {
		selectedBefore.Next = nil
		return
	}
	selectedBefore.Next = selectedCurr.Next
}

func (l *List) Push(i *Item){
	if l.Head != nil{
		i.Next = l.Head
	}
	
	l.Head = i
	l.Size = l.Size + 1
}

func (l *List) Pop() (*Item){
	old := l.Head
	l.Head = l.Head.Next	
	return old;
}