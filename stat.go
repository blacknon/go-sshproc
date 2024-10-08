// Copyright (c) 2024 Blacknon. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package sshproc

import (
	"io"
	"strconv"
	"strings"
	"time"

	proc "github.com/c9s/goprocinfo/linux"
)

func createCPUStat(fields []string) *proc.CPUStat {
	s := proc.CPUStat{}
	s.Id = fields[0]

	for i := 1; i < len(fields); i++ {
		v, _ := strconv.ParseUint(fields[i], 10, 64)
		switch i {
		case 1:
			s.User = v
		case 2:
			s.Nice = v
		case 3:
			s.System = v
		case 4:
			s.Idle = v
		case 5:
			s.IOWait = v
		case 6:
			s.IRQ = v
		case 7:
			s.SoftIRQ = v
		case 8:
			s.Steal = v
		case 9:
			s.Guest = v
		case 10:
			s.GuestNice = v
		}
	}
	return &s
}

func (p *ConnectWithProc) ReadStat(path string) (*proc.Stat, error) {
	file, err := p.sftp.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	content := string(b)
	lines := strings.Split(content, "\n")

	var stat proc.Stat = proc.Stat{}

	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if fields[0][:3] == "cpu" {
			if cpuStat := createCPUStat(fields); cpuStat != nil {
				if i == 0 {
					stat.CPUStatAll = *cpuStat
				} else {
					stat.CPUStats = append(stat.CPUStats, *cpuStat)
				}
			}
		} else if fields[0] == "intr" {
			stat.Interrupts, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "ctxt" {
			stat.ContextSwitches, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "btime" {
			seconds, _ := strconv.ParseInt(fields[1], 10, 64)
			stat.BootTime = time.Unix(seconds, 0)
		} else if fields[0] == "processes" {
			stat.Processes, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "procs_running" {
			stat.ProcsRunning, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "procs_blocked" {
			stat.ProcsBlocked, _ = strconv.ParseUint(fields[1], 10, 64)
		}
	}
	return &stat, nil
}
