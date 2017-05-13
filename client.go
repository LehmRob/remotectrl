// Copyright (c) 2017 Robert Lehmann <lehmrob@gmail.com>
// This software is licensed under BSD Clause 2 License
// Look into the LICENSE file for more information

package remotectrl

import (
	"io"
	"os/exec"

	"golang.org/x/crypto/ssh"
)

// Client holds information which are necessary to run the client
type Client struct {
	// Host is the address of the server
	Host string
	// Port is the port to connect on the server
	Port string
	// User which is used for authentification
	User string
	// Password is used for authentificate the user
	Pass string
}

func (c *Client) Run() error {
	sshConfig := &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Pass),
		},
	}

	conn, err := ssh.Dial("tcp", c.Host+c.Port, sshConfig)
	if err != nil {
		return err
	}

	session, err := conn.NewSession()
	if err != nil {
		return err
	}

	rStdin, err := session.StdinPipe()
	if err != nil {
		return err
	}

	rStdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	cmd := exec.Command("bash")
	lStdin, err := cmd.StdinPipe()
	if err != nil {
		return nil
	}

	lStdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.Run()

	go io.Copy(rStdin, lStdout)
	io.Copy(lStdin, rStdout)

	return nil
}
