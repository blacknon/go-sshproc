// Copyright(c) 2024 Blacknon. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/blacknon/go-sshlib"
	"github.com/blacknon/go-sshproc"

	"golang.org/x/crypto/ssh"
)

var (
	host     = "server1.example.com"
	port     = "22"
	user     = "user"
	key      = "~/.ssh/id_rsa"
	password = ""
)

func main() {
	// Create sshlib.Connect
	con := &sshlib.Connect{}

	// Create ssh.AuthMethod
	authMethod, _ := sshlib.CreateAuthMethodPublicKey(key, password)

	// Connect ssh server
	con.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	err := con.CreateClient(host, port, user, []ssh.AuthMethod{authMethod})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create sshproc
	proc := &sshproc.ConnectWithProc{Connect: con}
	err = proc.CreateSftpClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer proc.CloseSftpClient()

	fmt.Println("Read net/unix")
	stat, err := proc.ReadNetUnixDomainSockets("/proc/net/unix")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(stat)
}
