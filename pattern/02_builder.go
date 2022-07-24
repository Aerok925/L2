package main

import "fmt"

type ComputerBuilderI interface {
	CPU(string) ComputerBuilderI
	MB(string) ComputerBuilderI
	RAM(int) ComputerBuilderI

	Build() Computer
}

type officeComputerBuilder struct {
	ComputerBuilder
}

func (b officeComputerBuilder) Build() Computer {
	return Computer{
		cpu: b.cpu,
		mb:  b.mb,
		ram: b.ram,
	}
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return officeComputerBuilder{}.RAM(123).CPU("Хуета").MB("полная")
}

type Computer struct {
	cpu string
	mb  string
	ram int
}

type ComputerBuilder struct {
	cpu string
	mb  string
	ram int
}

func NewComputerBuilder() ComputerBuilderI {
	return ComputerBuilder{}
}

func (b ComputerBuilder) CPU(str string) ComputerBuilderI {
	b.cpu = str
	return b
}
func (b ComputerBuilder) MB(str string) ComputerBuilderI {
	b.mb = str
	return b
}
func (b ComputerBuilder) RAM(ram int) ComputerBuilderI {
	b.ram = ram
	return b
}
func (b ComputerBuilder) Build() Computer {
	return Computer{
		cpu: b.cpu,
		mb:  b.mb,
		ram: b.ram,
	}
}

type Director struct {
	b ComputerBuilderI
}

func NewDirector(b ComputerBuilderI) *Director {
	return &Director{
		b: b,
	}
}

func (d Director) BuildComputer() Computer {
	return d.b.Build()
}

func main() {
	//officeBuilder := NewOfficeComputerBuilder()
	compbuild := NewComputerBuilder().MB("qwe").CPU("Intel I5").RAM(4)
	direct := NewDirector(compbuild)
	fmt.Println(direct.BuildComputer())
}
