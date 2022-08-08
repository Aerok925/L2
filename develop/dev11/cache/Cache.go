package cache

import (
	_ "../storageI"
	"./LRU"
	"./cell"
	"encoding/json"
	"errors"
)

func New() *Cache {
	return &Cache{arr: LRU.New(100)}
}

type Cache struct {
	arr *LRU.LRU
}

func (r *Cache) Print() {
	r.arr.Print()
}

func (c *Cache) LoadIn(data *cell.Cell) error {
	c.arr.Append(*data)
	return nil
}

func (c *Cache) ConvertJSON(uuid string) ([]byte, error) {
	elem := c.arr.Get(uuid)
	if elem == nil {
		return nil, errors.New("Qwe")
	}
	ret, err := json.Marshal(elem)
	if err != nil && elem != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Cache) Update(data *cell.Cell) error {
	temp1 := c.arr.Get(data.Uuid)
	if temp1 == nil {
		return errors.New("Not found!")
	}
	c.arr.Update(*data)
	return nil
}

func (c *Cache) Delete(uuid, date string) error {
	return c.arr.Delete(uuid, date)
}
