package main

import (
	"fmt"
	"math"
)

/*
Посититель - это порождающий паттернб который позволяет выполнять операции с каждым объектом из некоторой структуры,
он позволяет определить новую операцию, не изменяя класс объекта
*/

type Visitor interface {
	visitorForSquare(*Square)
	visitorForCircle(*Circle)
	visitorForrectangle(*Rectangle)
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitorForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitorForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitorForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type TestVisitorGetData struct {
	test int
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitorForSquare(s *Square) {
	a.area = s.side * s.side
	fmt.Println("Calculating area for square", a.area)
}

func (a *AreaCalculator) visitorForCircle(s *Circle) {
	a.area = int(math.Pi * float64(s.radius*s.radius))
	fmt.Println("Calculating area for circle", a.area)
}

func (a *AreaCalculator) visitorForrectangle(s *Rectangle) {
	a.area = s.l * s.b
	fmt.Println("Calculating area for rectangle", a.area)
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}
	sh := make([]Shape, 0, 3)
	sh = append(sh, square)
	sh = append(sh, circle)
	sh = append(sh, rectangle)
	areaCalculator := &AreaCalculator{}

	for _, s := range sh {
		s.accept(areaCalculator)
	}
}
