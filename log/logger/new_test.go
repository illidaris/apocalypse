package logger

import "testing"

func TestNew(t *testing.T) {
	New(nil)
	Info("new log")
}

func ExampleNew() {
	New(nil)
}
