// Copyright (c) 2024 Blacknon. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package sshproc

import (
	"github.com/blacknon/go-sshlib"
	"github.com/pkg/sftp"
)

type ConnectWithProc struct {
	*sshlib.Connect
	sftp *sftp.Client
}

func (c *ConnectWithProc) CreateSftpClient() error {
	sftp, err := sftp.NewClient(c.Client)

	if err == nil {
		c.sftp = sftp
	}

	return err
}

func (c *ConnectWithProc) CloseSftpClient() error {
	return c.sftp.Close()
}
