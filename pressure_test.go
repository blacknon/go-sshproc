package sshproc

import "testing"

func TestParsePressureData(t *testing.T) {
	t.Parallel()

	data := []byte("some avg10=0.50 avg60=0.60 avg300=0.70 total=123\nfull avg10=0.05 avg60=0.06 avg300=0.07 total=45\n")

	stat, err := parsePressureData(data)
	if err != nil {
		t.Fatalf("parsePressureData returned error: %v", err)
	}

	if stat.Some == nil || stat.Full == nil {
		t.Fatalf("expected some/full metrics, got %+v", stat)
	}

	if stat.Some.Avg10 != 0.50 || stat.Some.Total != 123 {
		t.Fatalf("unexpected some metrics: %+v", stat.Some)
	}

	if stat.Full.Avg300 != 0.07 || stat.Full.Total != 45 {
		t.Fatalf("unexpected full metrics: %+v", stat.Full)
	}
}

func TestParsePressureDataSomeOnly(t *testing.T) {
	t.Parallel()

	data := []byte("some avg10=1.00 avg60=2.00 avg300=3.00 total=999\n")

	stat, err := parsePressureData(data)
	if err != nil {
		t.Fatalf("parsePressureData returned error: %v", err)
	}

	if stat.Some == nil {
		t.Fatalf("expected some metrics, got %+v", stat)
	}
	if stat.Full != nil {
		t.Fatalf("expected full metrics to be nil, got %+v", stat.Full)
	}
}
