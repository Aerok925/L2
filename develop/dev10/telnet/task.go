package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type telnet struct {
	host, port string
	timeout    time.Duration
}

func NewTelnet() *telnet {
	return &telnet{}
}

func (t *telnet) Parsing() bool {
	if len(os.Args) == 3 {
		t.host = os.Args[1]
		t.port = os.Args[2]
		return true
	}
	if len(os.Args) == 4 {
		arg := os.Args[1]
		substr := "--timeout="
		if strings.Contains(arg, substr) {
			timeDuration := strings.TrimPrefix(arg, substr)
			timeout, err := time.ParseDuration(timeDuration)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			t.timeout = timeout
		} else {
			return false
		}
		t.host = os.Args[2]
		t.port = os.Args[3]
		return true
	}
	return false
}

func (t *telnet) Connect(conn net.Conn, exit chan os.Signal) {
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	for {
		fmt.Print("from client: ")
		text, err := console.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading string: %v\n", err)
			exit <- syscall.SIGQUIT
			return
		}
		text = strings.TrimSpace(text)

		fmt.Fprintf(conn, text+"\n")
		if text == "exit" {
			fmt.Fprintf(os.Stdout, "%s\n", "connection closed")
			exit <- syscall.SIGQUIT
			return
		}

		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from conn: %v\n", err)
			exit <- syscall.SIGQUIT
			os.Exit(1)
		}
		message = strings.TrimSpace(message)
		fmt.Printf("from server: %s\n", message)
		log.Println("lol")
	}
}

func main() {
	t := NewTelnet()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)

	ok := t.Parsing()
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", "pizda")
		os.Exit(1)
	}
	d := net.Dialer{Timeout: t.timeout}
	conn, err := d.Dial("tcp", t.host+":"+t.port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	go t.Connect(conn, quit)
	select {
	case <-quit:
		fmt.Fprintf(os.Stdout, "%s\n", "programm finished")
		os.Exit(0)
	}
}
