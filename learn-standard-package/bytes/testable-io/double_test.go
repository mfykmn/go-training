package main

import (
	"bytes"
	"testing"
)

func TestDouble(t *testing.T) {
	stdin := bytes.NewBufferString("foo\n")
	stdout := new(bytes.Buffer)

	err := Double(stdin, stdout)
	if err != nil {
		t.Fatal("failed to call Double(): %s", err)
	}

	expected := []byte("foo\nfoo\n")

	if bytes.Compare(expected, stdout.Bytes()) != 0 {
		t.Fatal("not matched")
	}
}
