//go:build linux

package gocapng

import (
	"testing"
)

func TestInitFunctions(t *testing.T) {
	caps, err := InitFunctions()
	if err != nil {
		t.Error(err)
	}
	if caps == nil {
		t.Error("caps is nil")
	}
	caps = nil
}
