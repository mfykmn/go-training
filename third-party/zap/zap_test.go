package zap

import (
	"testing"
)

func init() {
	New(Info)
}

func TestGetLogger(t *testing.T) {
	Fatalf("trace00000001", "Failed DB connection. errCode is %s. %s", "0001", "TIMEOUT")
}
