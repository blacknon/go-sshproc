package sshproc

import (
	"testing"

	proc "github.com/c9s/goprocinfo/linux"
)

func TestCountTCPSocketStates(t *testing.T) {
	t.Parallel()

	sockets := &proc.NetTCPSockets{
		Sockets: []proc.NetTCPSocket{
			{NetSocket: proc.NetSocket{Status: 0x01}},
			{NetSocket: proc.NetSocket{Status: 0x01}},
			{NetSocket: proc.NetSocket{Status: 0x0A}},
			{NetSocket: proc.NetSocket{Status: 0xFF}},
		},
	}

	counts := CountTCPSocketStates(sockets)

	if counts["ESTABLISHED"] != 2 {
		t.Fatalf("unexpected ESTABLISHED count: %d", counts["ESTABLISHED"])
	}
	if counts["LISTEN"] != 1 {
		t.Fatalf("unexpected LISTEN count: %d", counts["LISTEN"])
	}
	if counts["UNKNOWN_FF"] != 1 {
		t.Fatalf("unexpected UNKNOWN_FF count: %d", counts["UNKNOWN_FF"])
	}
}

func TestCountUDPSocketStates(t *testing.T) {
	t.Parallel()

	sockets := &proc.NetUDPSockets{
		Sockets: []proc.NetUDPSocket{
			{NetSocket: proc.NetSocket{Status: 0x07}},
			{NetSocket: proc.NetSocket{Status: 0x07}},
			{NetSocket: proc.NetSocket{Status: 0x0A}},
		},
	}

	counts := CountUDPSocketStates(sockets)

	if counts["CLOSE"] != 2 {
		t.Fatalf("unexpected CLOSE count: %d", counts["CLOSE"])
	}
	if counts["LISTEN"] != 1 {
		t.Fatalf("unexpected LISTEN count: %d", counts["LISTEN"])
	}
}
