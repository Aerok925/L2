package main

import "fmt"

// Цепь вызовов - это поведенческий паттерн, котороый позволяет передовать запросы последовательно по цепи обработчиков

// Departament - это интерфейс нашего обработчика, в данном случае он имеет метод запуска и установки следующего за ним обработчика
type Departament interface {
	execute(Patient)
	setNext(Departament)
}

// Reсeption - один оз обработчиков, который реализует интерфейс Departament
type Reсeption struct {
	next Departament
}

func (r *Reсeption) execute(p Patient) {
	if p.GetProblem() == "Ухо" {
		p.SetStatus(0)
		fmt.Println("Вам следует пойти к Отоларингологу")
	} else if p.GetProblem() == "Перелом" {
		p.SetStatus(1)
		fmt.Println("Вам следует пойти к хирургу")
	} else {
		p.SetStatus(2)
		fmt.Println("хм, непонятно что с вами, следует пойти терапевту")
	}
	r.next.execute(p)
	return
}

func (r *Reсeption) setNext(dep Departament) {
	r.next = dep
}

type Otolaryngologist struct {
	next Departament
}

func (o *Otolaryngologist) execute(p Patient) {
	if p.GetStatus() == 0 {
		fmt.Println("Да, сейчас разберемся с ваших ухом")
	}
	o.next.execute(p)
}

func (o *Otolaryngologist) setNext(dep Departament) {
	o.next = dep
}

type Surgeon struct {
	next Departament
}

func (s *Surgeon) execute(p Patient) {
	if p.GetStatus() == 1 {
		fmt.Println("Сейчас наложим гипс и норм будет")
	}
	s.next.execute(p)
}

func (s *Surgeon) setNext(dep Departament) {
	s.next = dep
}

type Therapist struct {
	next Departament
}

func (t *Therapist) execute(p Patient) {
	if p.GetStatus() == 2 {
		fmt.Println("Сейчас разберемся с вашей проблемой")
	}
	t.next.execute(p)
}

func (t *Therapist) setNext(dep Departament) {
	t.next = dep
}

type Pharmacy struct {
	next Departament
}

func (ph *Pharmacy) execute(p Patient) {
	fmt.Println("Вот выписанные доктором лекарства, проходите на кассу")
	ph.next.execute(p)
}

func (ph *Pharmacy) setNext(d Departament) {
	ph.next = d
}

type CashRegister struct {
	next Departament
}

func (cash *CashRegister) execute(p Patient) {
	fmt.Println("Спасибо за покупку!!")
	cash.next.execute(p)
}

func (cash *CashRegister) setNext(dep Departament) {
	cash.next = dep
}

// Затычка цепи, которая ничего не делает, обязана стоять в конце цепи
type Plug struct {
}

func (p *Plug) execute(Patient) {

}

func (p *Plug) setNext(Departament) {

}

// Patient - это интерфейс нашего запроса, который будет передаваться по цепи вызовов
type Patient interface {
	SetProblem(string)
	GetProblem() string
	SetStatus(int)
	GetStatus() int
}

// invalid - вструктура реализующая интерфейс Patient, у больного есть проблема, исходя из неё, пациент будет направлен
// к одному из специалистов
type invalid struct {
	problem string
	status  int
}

func (inv *invalid) SetProblem(str string) {
	inv.problem = str
}

func (inv *invalid) GetProblem() string {
	return inv.problem
}

func (inv *invalid) SetStatus(status int) {
	inv.status = status
}

func (inv *invalid) GetStatus() int {
	return inv.status
}

// Конструктор структуры invalid
func NewInvalid() Patient {
	return &invalid{}
}

func main() {
	// создаем новые объекты
	hospital := &Reсeption{}
	otolar := &Otolaryngologist{}
	surgeon := &Surgeon{}
	therapist := &Therapist{}
	pharmacy := &Pharmacy{}
	cash := &CashRegister{}
	plug := &Plug{}

	// устанавливаем связь меж ними
	hospital.setNext(otolar)
	otolar.setNext(surgeon)
	surgeon.setNext(therapist)
	therapist.setNext(pharmacy)
	pharmacy.setNext(cash)
	cash.setNext(plug)

	inv := NewInvalid()

	inv.SetProblem("жесть полная")
	hospital.execute(inv)

	inv.SetProblem("Ухо")
	hospital.execute(inv)

	inv.SetProblem("Перелом")
	hospital.execute(inv)
}
