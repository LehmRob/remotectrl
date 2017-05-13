package remotectrl

import (
	"fmt"
	"io"

	"github.com/gliderlabs/ssh"
)

// ServerConfig holds information for the remotectrl server
type Server struct {
	// Port is the port which will be used fot the server to listen
	Port string
	// User for the registration on the remotectrl server
	User string
	// Pass is the password phrase for the user to login on the server
	Pass string
}

func (s *Server) Run() error {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello World %s\n", s.User()))
	})

	return ssh.ListenAndServe(s.Port, nil)
}
