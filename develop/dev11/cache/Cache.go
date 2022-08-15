package cache

import (
	"encoding/json"
	"errors"
	"github.com/Aerok925/L2/develop/dev11/cache/LRU"
	"github.com/Aerok925/L2/develop/dev11/cache/cell"
)

func New() *Cache {
	return &Cache{arr: LRU.New(100)}
}

var Storage *Cache = New()

type Cache struct {
	arr *LRU.LRU
}

func (r *Cache) Print() {
	r.arr.Print()
}

func (r *Cache) Get(key string) []cell.Cell {
	return r.arr.Get(key)
}

func (c *Cache) LoadIn(data *cell.Cell) error {
	if c.arr.Append(*data) {
		return nil
	}
	return errors.New("Can`t append!")
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
	if c.arr.Update(*data) {
		return nil
	}
	return errors.New("Not found!")
}

func (c *Cache) Delete(data *cell.Cell) error {
	return c.arr.Delete(data.Uuid, data.Date)
}
