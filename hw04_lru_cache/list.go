package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
	}

	if l.front != nil {
		l.front.Prev = item
		item.Next = l.front
	}

	if l.back == nil {
		l.back = item
	}

	l.front = item

	l.len++

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
	}

	if l.back != nil {
		l.back.Next = item
		item.Prev = l.back
	}

	if l.front == nil {
		l.front = item
	}

	l.back = item

	l.len++

	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
