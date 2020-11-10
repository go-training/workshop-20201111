package main

import "log"

type option func(*smtp)

func withAddress(address string) option {
	return func(s *smtp) {
		s.Address = address
	}
}

func withPort(port string) option {
	return func(s *smtp) {
		s.Port = port
	}
}

type smtp struct {
	Address string
	Port    string
}

func (s *smtp) send() {
	log.Println("send email from:", s.Address, "port:", s.Port)
}

func new(opts ...option) *smtp {
	s := &smtp{
		Address: "127.0.0.1",
		Port:    "25",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func main() {
	foo := new()
	foo.send()
	bar := new(withAddress("172.168.1.1"))
	bar.send()
}
