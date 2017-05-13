package main

import (
	"flag"
	"log"

	"github.com/lehmrob/remotectrl"
)

var (
	port   = flag.String("port", ":2020", "Port for the server for listening")
	user   = flag.String("user", "user", "User for the server to register")
	pass   = flag.String("pass", "secret", "Password which is needed to register on the server")
	server = flag.Bool("server", false, "Flag indicates if the daemon should work in server mode or not")
	host   = flag.String("host", "", "Host is the host address for the server, use this flag only in client mode (no -server flag)")
)

func main() {
	log.Println("Hello World")
	flag.Parse()

	if *server {
		log.Println("Starting Server")
		s := remotectrl.Server{
			Port: *port,
			User: *user,
			Pass: *pass,
		}

		log.Fatal(s.Run())
	} else {
		log.Println("Starting Client")
		c := remotectrl.Client{
			Host: *host,
			Port: *port,
			User: *user,
			Pass: *pass,
		}
		log.Fatal(c.Run())
	}
}
