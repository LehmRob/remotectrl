// Copyright (c) 2017 Robert Lehmann <lehmrob@gmail.com>
// This software is licensed under BSD Clause 2 License
// Look into the LICENSE file for more information

package remotectrl

import (
	"io"
	"os"

	"github.com/gliderlabs/ssh"
)

// Server holds information for the remotectrl server
type Server struct {
	// Port is the port which will be used fot the server to listen
	Port string
	// User for the registration on the remotectrl server
	User string
	// Pass is the password phrase for the user to login on the server
	Pass string

	server ssh.Server
}

func (s *Server) Run() error {
	s.server = ssh.Server{
		Addr: s.Port,
		PasswordHandler: func(ctx ssh.Context, pass string) bool {
			return ctx.User() == s.User && pass == s.Pass
		},
	}

	s.server.Handle(func(s ssh.Session) {
		for {
			go io.Copy(s, os.Stdout)
			io.Copy(os.Stdin, s)
		}
	})

	return s.server.ListenAndServe()
}
