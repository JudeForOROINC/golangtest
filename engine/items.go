package engine

import (
	"container/list"
	"errors"
	"fmt"
)

type Item struct {
	id            int
	name          string
	value         int
	weight        int
	owner         *Items
	ownersElement *list.Element
}

type Items struct {
	items list.List
}

func (items *Items) AddItem(i *Item) (Items, error) {
	// items.items = append(items.items, i)
	e := items.items.PushBack(i)
	i.owner = items
	i.ownersElement = e
	return *items, nil
}

func (items *Items) DeleteItem(i *Item) (Items, error) {
	// items.items = append(items.items, i)
	if i.ownersElement == nil {
		return *items, errors.New("no Item found")
	}
	items.items.Remove(i.ownersElement)
	i.owner = nil
	i.ownersElement = nil
	return *items, nil
}

func NewItem(id int, name string, weight int, value int) *Item {
	item := Item{}
	item.name = name
	item.id = id
	item.weight = weight
	item.value = value

	return &item
}

func (item *Item) MoveTo(items *Items) {
	if item.owner != nil {
		item.owner.DeleteItem(item)
	}

	items.AddItem(item)
}

func Main() {
	fmt.Println("new module")
}
