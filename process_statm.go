// Copyright (c) 2024 Blacknon. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package sshproc

import (
	"io"
	"strconv"
	"strings"

	proc "github.com/c9s/goprocinfo/linux"
)

func (p *ConnectWithProc) ReadProcessStatm(path string) (*proc.ProcessStatm, error) {
	file, err := p.sftp.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	s := string(b)
	f := strings.Fields(s)

	statm := proc.ProcessStatm{}

	var n uint64

	for i := 0; i < len(f); i++ {

		if n, err = strconv.ParseUint(f[i], 10, 64); err != nil {
			return nil, err
		}

		switch i {
		case 0:
			statm.Size = n
		case 1:
			statm.Resident = n
		case 2:
			statm.Share = n
		case 3:
			statm.Text = n
		case 4:
			statm.Lib = n
		case 5:
			statm.Data = n
		case 6:
			statm.Dirty = n
		}

	}

	return &statm, nil
}
