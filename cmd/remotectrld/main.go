package main

import (
	"flag"
	"log"
)

var (
	server = flag.String("port", ":2020", "Port for the server for listening")
	user   = flag.String("user", "user", "User for the server to register")
	pass   = flag.String("pass", "secret", "Password which is needed to register on the server")
)

func main() {
	log.Println("Hello World")
	flag.Parse()
}
