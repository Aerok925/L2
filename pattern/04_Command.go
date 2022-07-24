package main

import "fmt"

// Интерфейс для любого устройства
type Device interface {
	On()
	Off()
	Call()
}

// Интерфейс для команды
type Command interface {
	Execute()
}

type CommandCall struct {
	device Device
}

func (com *CommandCall) Execute() {
	com.device.Call()
}

// Команда включения
type CommandOn struct {
	device Device
}

func (com *CommandOn) Execute() {
	com.device.On()
}

// Команда выключения
type CommandOff struct {
	device Device
}

func (com *CommandOff) Execute() {
	com.device.Off()
}

// Структура реализации любой кнопки
type Button struct {
	command Command
}

// метод для исполения реализации кнопки
func (b *Button) press() {
	b.command.Execute()
}

func (b *Button) SetCommand(com Command) {
	b.command = com
}

/// Структура реализующая интерфейс телефона
type Phone struct {
	IsWork bool
}

func (p *Phone) On() {
	p.IsWork = true
	fmt.Println("Phone on!")
}
func (p *Phone) Off() {
	p.IsWork = false
	fmt.Println("Phone off!")
}

func (p *Phone) Call() {
	fmt.Println("Call the number")
}

func NewPhone() *Phone {
	return &Phone{IsWork: false}
}

func NewButton(com Command) *Button {
	return &Button{command: com}
}

func main() {
	phone := NewPhone()
	commandOn := &CommandOn{device: phone}
	commandOff := &CommandOff{device: phone}
	commandCall := &CommandCall{device: phone}
	button := NewButton(commandOn)
	button.press()
	button.SetCommand(commandCall)
	button.press()
	button.SetCommand(commandOff)
	button.press()
}
