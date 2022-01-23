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

func TestCapabilityToName(t *testing.T) {
	capToName := map[Capability]string{
		CAPCHOWN:             "chown",
		CAPDACOverride:       "dac_override",
		CAPDACReadSearch:     "dac_read_search",
		CAPFOwner:            "fowner",
		CAPCheckpointRestore: "checkpoint_restore",
	}

	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	for cap, name := range capToName {
		s := caps.CapabilityToName(cap)
		if s != name {
			t.Errorf("Expected %s, got %s", name, s)
		}
	}
}

func TestEmptyCapabilityToName(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	s := caps.CapabilityToName(CAPCheckpointRestore + 1000)
	if s != "" {
		t.Errorf("expected s to empty, got %s instead", s)
	}
}

func TestNameToCapabilityFound(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	cap, err := caps.NameToCapability("dac_override")
	if err != nil {
		t.Errorf("Expected cap, but got error: %s", err)
	}

	if cap != CAPDACOverride {
		t.Errorf("Expected %d but got %d", CAPDACOverride, cap)
	}
}

func TestNameToCapabilityNotFound(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	cap, err := caps.NameToCapability("dacoverride")
	if err == nil {
		t.Errorf("expected err to be %s, but got nil", ErrCapabilityNotFound)
	}

	if cap != 0 {
		t.Errorf("Expected cap to be 0, but got %d", cap)
	}
}

func TestPrintCapsTextEmpty(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	buf := caps.PrintCapsText(
		PrintBuffer,
		TypeInheritable,
	)

	if buf != "" {
		t.Errorf("Expected buf to be '' but '%s' found", buf)
	}

	buf = caps.PrintCapsText(
		PrintStdOut,
		TypeInheritable,
	)

	if buf != "" {
		t.Errorf("buf expected to be empty but '%s' found", buf)
	}
}

func TestPrintCapsTextBufValid(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	caps.Clear(SelectAll)

	buf := caps.PrintCapsText(
		PrintBuffer,
		TypeInheritable,
	)

	if buf != "none" {
		t.Errorf("Expected buf to be 'none' but '%s' found", buf)
	}

}
