package doublylinkedlist

import (
	"fmt"

	"github.com/pkg/errors"
)

type List struct {
	Lenght    int
	FirstElem *Item
	LastElem  *Item
}

func NewList() *List {
	return &List{}
}

func (l *List) Len() int {
	return l.Lenght
}

func (l *List) First() *Item {
	return l.FirstElem
}

func (l *List) Last() *Item {
	return l.LastElem
}

func (l *List) PushFront(v interface{}) {
	if l.LastElem == nil {
		l.LastElem = NewItem(v)
		l.FirstElem = l.LastElem
		l.Lenght = 1
	} else {
		newItem := NewItem(v)
		newItem.PrevElem = l.LastElem
		l.LastElem.NextElem = newItem
		l.LastElem = newItem
		l.Lenght++
	}
}

func (l *List) PushBack(v interface{}) {
	if l.FirstElem == nil {
		l.FirstElem = NewItem(v)
		l.LastElem = l.FirstElem
		l.Lenght = 1
	} else {
		newItem := NewItem(v)
		newItem.NextElem = l.FirstElem
		l.FirstElem.PrevElem = newItem
		l.FirstElem = newItem
		l.Lenght++
	}
}

func (l *List) Remove(i Item) {
	if l.Len() == 0 {
		return
	} else if l.Len() == 1 {
		l.FirstElem = nil
		l.LastElem = nil
	}

	item := l.FirstElem
	if item.NextElem == i.NextElem && item.Val == i.Val && item.PrevElem == i.PrevElem {
		if item.NextElem != nil && item.PrevElem != nil {
			item.PrevElem.NextElem = item.NextElem
		} else if item.PrevElem != nil {
			item.PrevElem.NextElem = nil
		}

		if item.PrevElem != nil && item.NextElem != nil {
			item.NextElem.PrevElem = item.PrevElem
		} else if item.NextElem != nil {
			item.NextElem.PrevElem = nil
		}

		l.Lenght--
		return
	}
	for item.Next() != nil {
		item = item.Next()
		if item.NextElem == i.NextElem && item.Val == i.Val && item.PrevElem == i.PrevElem {
			item.PrevElem.NextElem = item.NextElem
			item.NextElem.PrevElem = item.PrevElem
			l.Lenght--
			break
		}
	}
}

type Item struct {
	Val      interface{}
	PrevElem *Item
	NextElem *Item
}

func NewItem(v interface{}) *Item {
	return &Item{Val: v}
}

func (i *Item) Next() *Item {
	if i.NextElem == nil {
		errors.Errorf("Next elem is not exist")
		return nil
	}
	return i.NextElem
}

func (i *Item) Prev() *Item {
	if i.PrevElem == nil {
		fmt.Println(errors.Errorf("Prev elem is not exist"))
		return nil
	}
	return i.PrevElem
}

func (i *Item) Value() interface{} {
	if i == nil {
		fmt.Println(errors.Errorf("Elem is not exist"))
		return nil
	}
	return i.Val
}
