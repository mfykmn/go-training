package log

import (
	"testing"
)

func init() {
	New(Info)
}

func TestGetLogger(t *testing.T) {
	Fatalf("aaaaaa %s", "bbbb")
}
