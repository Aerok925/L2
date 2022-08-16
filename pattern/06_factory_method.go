package main

import (
	"errors"
	"fmt"
)

// Паттерн Фабрика метод- это порождающий паттерн, который определяет интерфейс, который реализует супер класс
// в свою очерь суперкласс наследуется дочерним объектам, что позволяет изменять тип создоваемых объектов

// transport - интерфейс, которому должен соответствовать супер класс
type transport interface {
	delivery()
	SetCity(string)
	SetShipment(string)
	GetShipment() string
	GetSpeed() int
}

// LandTransport - суперкласс, от которого будут наследовать весь наземный транспорт
type LandTransport struct {
	shipment string
	city     string
	speed    int
}

func (land *LandTransport) delivery() {
	fmt.Println("Поехал доставлять", land.shipment, "в", land.city)
}
func (land *LandTransport) SetCity(city string) {
	land.city = city
}
func (land *LandTransport) SetShipment(shipment string) {
	land.shipment = shipment
}
func (land *LandTransport) GetShipment() string {
	return land.shipment
}
func (land *LandTransport) GetSpeed() int {
	return land.speed
}

// WaterTransport - суперкласс, от которого будут наследовать весь водный транспорт
type WaterTransport struct {
	shipment string
	city     string
	speed    int
}

func (water *WaterTransport) delivery() {
	fmt.Println("Поплыл доставлять", water.shipment, "в", water.city)
}
func (water *WaterTransport) SetCity(city string) {
	water.city = city
}
func (water *WaterTransport) SetShipment(shipment string) {
	water.shipment = shipment
}
func (water *WaterTransport) GetShipment() string {
	return water.shipment
}
func (water *WaterTransport) GetSpeed() int {
	return water.speed
}

// Zhiguli - дочерний классб который наследуется от LandTransport
type Zhiguli struct {
	LandTransport
}

// FishingBoat - дочерний класс, который наследуется от WaterTransport
type FishingBoat struct {
	WaterTransport
}

// Конструкторы
func NewZhiguli() transport {
	return &Zhiguli{LandTransport{speed: 5}}
}

func NewFishingBoat() transport {
	return &FishingBoat{WaterTransport: WaterTransport{speed: 3}}
}

// CreateTransport - функция, которая создает объект различных типов
func CreateTransport(trans string) (transport, error) {
	switch trans {
	case "Zhiquli":
		return NewZhiguli(), nil
	case "FishingBoat":
		return NewFishingBoat(), nil
	default:
		return nil, errors.New("Don`t found this transport")
	}
}

func main() {
	car, _ := CreateTransport("Zhiquli")
	car.SetShipment("товар")
	car.SetCity("Култаево")
	car.delivery()
	boat, _ := CreateTransport("FishingBoat")
	boat.SetShipment("телефон")
	boat.SetCity("Мск")
	boat.delivery()
}
