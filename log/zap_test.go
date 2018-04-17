package log

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	New()
	logger := GetLogger()

	logger.Fatalf("aaaaaa %s", "bbbb")
}
