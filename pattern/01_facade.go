package main

import (
	"fmt"
)

/*
Фасад является структурным паттерном, который предоставляет единственный интерфейс, вместо набора интерфейсов
нокоторой подсистемы. Другими словами он определяет интерфейс более высого уровня, который упрощает пользователю общение
с подсистемой.
*/

type Courier interface {
	Delivery(string)
	PickUpProducts(i StoreI)
	GetName() string
}

type DeliverMen struct {
	name string
}

func (d *DeliverMen) Delivery(address string) {
	fmt.Println(d.name, "доставил продукты до", address)
}

func (d *DeliverMen) GetName() string {
	return d.name
}

func NewDiliverMen() Courier {
	return &DeliverMen{name: "Вася"}
}

func (d *DeliverMen) PickUpProducts(i StoreI) {
	fmt.Println(d.GetName(), "забирает продукты в магазине", i.GetName())
	i.ProvideProduct(d)
}

type StoreI interface {
	ProvideProduct(courier Courier)
	GetName() string
}

type Magnit struct {
	name string
}

func (m *Magnit) GetName() string {
	return m.name
}

func (m *Magnit) ProvideProduct(c Courier) {
	fmt.Println(m.GetName(), "Выдал продукты курьеру", c.GetName())
}

func NewMagnit() StoreI {
	return &Magnit{name: "Magnit"}
}

type DelivEda interface {
	DeliveryTo(string)
}

type DelivertClub struct {
	name string
}

func (d *DelivertClub) DeliveryTo(name string) {
	cour := NewDiliverMen()
	shop := NewMagnit()
	cour.PickUpProducts(shop)
	cour.Delivery(name)
}

func NewDeliverClub() DelivEda {
	return &DelivertClub{name: "qwe"}
}

func main() {
	dev := NewDiliverMen()
	shop := NewMagnit()
	dev.PickUpProducts(shop)
	dev.Delivery("дома")

	ogr := NewDeliverClub()
	ogr.DeliveryTo("работы")
}
