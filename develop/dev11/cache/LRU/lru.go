package LRU

import (
	"../cell"
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Item struct {
	Key   string
	Value []cell.Cell
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	mutex    sync.RWMutex
	queue    *list.List
}

func New(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRU) Print() {
	for key, _ := range c.items {
		fmt.Println(key, c.Get(key))
	}
}

func (c *LRU) Append(value cell.Cell) bool {
	if element, exists := c.items[value.Uuid]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = append(element.Value.(*Item).Value, value)
		return true
	}
	if c.queue.Len() == c.capacity {
		c.purge()
	}
	item := &Item{
		Key:   value.Uuid,
		Value: make([]cell.Cell, 0),
	}
	item.Value = append(item.Value, value)
	element := c.queue.PushFront(item)
	c.items[item.Key] = element
	return true
}

func (c *LRU) Update(value cell.Cell) bool {
	if element, exists := c.items[value.Uuid]; exists == true {
		c.queue.MoveToFront(element)
		for i, elem := range element.Value.(*Item).Value {
			if value.Date == elem.Date {
				element.Value.(*Item).Value[i].Event = value.Event
				return true
			}
		}
	}
	return false
}

func (c *LRU) Delete(key string, date string) error {
	element, exists := c.items[key]
	if exists == false {
		return errors.New("Not found!")
	}
	for i, elem := range element.Value.(*Item).Value {
		if elem.Date == date {
			element.Value.(*Item).Value = append(element.Value.(*Item).Value[0:i], element.Value.(*Item).Value[i+1:]...)
		}
	}
	if len(element.Value.(*Item).Value) == 0 {
		c.queue.Remove(element)
		delete(c.items, key)
	}
	return nil
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}

func (c *LRU) Get(key string) []cell.Cell {
	element, exists := c.items[key]
	if exists == false {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}
