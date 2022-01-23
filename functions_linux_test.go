//go:build linux

package gocapng

import (
	"testing"
)

func TestInitFunctions(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	caps = nil
}
