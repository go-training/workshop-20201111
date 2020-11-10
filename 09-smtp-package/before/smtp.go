package main

import "log"

type smtp struct {
	Address string
	Port    string
}

func (s *smtp) send() {
	log.Println("send email")
}

func new(address, port string) *smtp {
	return &smtp{
		Address: address,
		Port:    port,
	}
}

func main() {
	foo := new("127.0.0.1", "25")
	foo.send()
}
