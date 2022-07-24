package main

import "fmt"

type flamethrower struct {
	charge int
}

func (f *flamethrower) Fire() {
	if f.charge-1 < 0 {
		fmt.Println("I can`t")
		return
	}
	f.charge--
	fmt.Println("Fire show!!!")
}

func (f *flamethrower) Charging() {
	f.charge = 10
}

type FireShow interface {
	Show()
	Reboot()
}

type Show struct {
	flame flamethrower
}

func (sh *Show) Show() {
	sh.flame.Fire()
}

func (sh *Show) Reboot() {
	sh.flame.Charging()
}

func NewFlamethrower(i int) *flamethrower {
	return &flamethrower{charge: i}
}

func NewShow(i int) *Show {
	return &Show{flame: *NewFlamethrower(i)}
}

func main() {
	show := NewShow(10)
	show.Show()
}
