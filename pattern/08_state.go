package main

import (
	"errors"
	"fmt"
)

type StateI interface {
	Connect() error
	Disconnect() error
	Send() error
}

type ListeningState struct {
	serv *Server
}

func (l *ListeningState) Connect() error {
	fmt.Println("Подключено")
	l.serv.SetState(l.serv.Established)
	return nil
}

func (l *ListeningState) Disconnect() error {
	fmt.Println("Подключения нет")
	return errors.New("Не было подключения")
}

func (l *ListeningState) Send() error {
	fmt.Println("Отправлять некуда")
	return errors.New("Отправлять некуда")
}

func NewListenState(s *Server) StateI {
	return &ListeningState{serv: s}
}

type ClosedState struct {
	serv *Server
}

func (l *ClosedState) Connect() error {
	fmt.Println("Закрыто")
	return errors.New("Закрыто")
}

func (l *ClosedState) Disconnect() error {
	fmt.Println("Закрыто")
	return errors.New("Закрыто")
}

func (l *ClosedState) Send() error {
	fmt.Println("Закрыто")
	return errors.New("Закрыто")
}

func NewCLosedState(s *Server) StateI {
	return &ClosedState{serv: s}
}

type EstablishedState struct {
	serv *Server
}

func (l *EstablishedState) Connect() error {
	fmt.Println("Уже есть")
	return errors.New("Уже есть")
}

func (l *EstablishedState) Disconnect() error {
	fmt.Println("Отключено")
	l.serv.SetState(l.serv.listen)
	return nil
}

func (l *EstablishedState) Send() error {
	fmt.Println("Сообщение отправлено")
	return nil
}

func NewEstablishedState(s *Server) StateI {
	return &EstablishedState{serv: s}
}

type Server struct {
	listen      StateI
	Close       StateI
	Established StateI

	currentState StateI
}

func (s *Server) SetState(st StateI) {
	s.currentState = st
}

func NewServer() *Server {
	retValue := &Server{}
	retValue.listen = NewListenState(retValue)
	retValue.Close = NewCLosedState(retValue)
	retValue.Established = NewEstablishedState(retValue)
	return retValue
}

func (s *Server) Start() {
	fmt.Println("Сервер запущен")
	s.SetState(s.listen)
}

func (s *Server) Closed() {
	fmt.Println("Сервер выключен")
	s.SetState(s.Close)
}

func (s *Server) Connect() error {
	return s.currentState.Connect()
}
func (s *Server) Disconnect() error {
	return s.currentState.Disconnect()
}
func (s *Server) Send() error {
	return s.currentState.Send()
}

func main() {
	serv := NewServer()
	serv.Start()
	err := serv.Connect()
	err = serv.Connect()
	err = serv.Disconnect()
	serv.Closed()
	err = serv.Connect()
	serv.Start()
	err = serv.Send()
	err = serv.Connect()
	err = serv.Send()
	fmt.Println(err)
}
