package sshproc

import "testing"

func TestParseUint64Data(t *testing.T) {
	t.Parallel()

	value, err := parseUint64Data([]byte("4096\n"))
	if err != nil {
		t.Fatalf("parseUint64Data returned error: %v", err)
	}

	if value != 4096 {
		t.Fatalf("unexpected value: %d", value)
	}
}

func TestParseFileNrData(t *testing.T) {
	t.Parallel()

	fileNr, err := parseFileNrData([]byte("1024\t0\t9223372036854775807\n"))
	if err != nil {
		t.Fatalf("parseFileNrData returned error: %v", err)
	}

	if fileNr.Allocated != 1024 || fileNr.Unused != 0 || fileNr.Max != 9223372036854775807 {
		t.Fatalf("unexpected file-nr: %+v", fileNr)
	}
}
