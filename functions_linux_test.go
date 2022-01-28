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

func TestUnInitializedFeatures(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	t.Run("testPrintCapsText", func(t2 *testing.T) {
		buf := caps.PrintCapsText(
			PrintBuffer,
			TypeInheritable,
		)

		if buf != "" {
			t2.Errorf("Expected buf to be '' but '%s' found", buf)
		}

		buf = caps.PrintCapsText(
			PrintStdOut,
			TypeInheritable,
		)

		if buf != "" {
			t2.Errorf("buf expected to be empty but '%s' found", buf)
		}

	})

	t.Run("testPrintCapsNumberic", func(t2 *testing.T) {
		buf := caps.PrintCapsNumberic(
			PrintBuffer,
			SelectAll,
		)
		if buf != "" {
			t2.Errorf("expected empty buf, but '%s' found", buf)
		}
	})

	t.Run("testHaveCapability", func(t2 *testing.T) {
		result := caps.HaveCapability(TypeInheritable, CAPCHOWN)
		if result {
			t2.Error("Expected result to be false, but it is true")
		}
	})

	t.Run("testHavePermittedCapabilities", func(t2 *testing.T) {
		result := caps.HavePermittedCapabilities()
		if result != ResultNone {
			t2.Errorf("Expected ResultNone, but %d (%s) found", result, result)
		}
	})

	t.Run("testHaveCapabilities", func(t2 *testing.T) {
		result := caps.HaveCapabilities(SelectAll)
		if result != ResultPartial {
			t2.Errorf("Expected ResultPartial, but %d (%s) found", result, result)
		}
	})
}

func TestInitializedFeatures(t *testing.T) {
	caps := Init()
	if caps == nil {
		t.Error("caps is nil")
	}
	defer func() {
		caps = nil
	}()

	caps.Clear(SelectAll)

	t.Run("testPrintCapsText", func(t2 *testing.T) {
		buf := caps.PrintCapsText(
			PrintBuffer,
			TypeInheritable,
		)

		if buf != "none" {
			t2.Errorf("Expected buf to be 'none' but '%s' found", buf)
		}

	})

	t.Run("testPrintCapsNumberic", func(t2 *testing.T) {
		buf := caps.PrintCapsNumberic(
			PrintBuffer,
			SelectAll,
		)

		expected := `Effective:   00000000, 00000000
Permitted:   00000000, 00000000
Inheritable: 00000000, 00000000
Bounding Set: 00000000, 00000000
Ambient Set: 00000000, 00000000
`
		if expected != buf {
			t2.Errorf("expected '%s' but found '%s'", expected, buf)
		}

	})
	t.Run("testHaveCapability", func(t2 *testing.T) {
		result := caps.HaveCapability(TypeInheritable, CAPCHOWN)
		if result {
			t2.Error("Expected result to be false, but it is true")
		}
	})

	t.Run("testHavePermittedCapabilities", func(t2 *testing.T) {
		result := caps.HavePermittedCapabilities()
		if result != ResultNone {
			t2.Errorf("Expected ResultNone, but %d (%s) found", result, result)
		}
	})

	t.Run("testHaveCapabilities", func(t2 *testing.T) {
		result := caps.HaveCapabilities(SelectAll)
		if result != ResultNone {
			t2.Errorf("Expected ResultNone, but %d (%s) found", result, result)
		}
	})

	t.Run("testFillAndHaveCapabilities", func(t2 *testing.T) {
		caps.Fill(SelectAll)
		result := caps.HaveCapabilities(SelectAll)
		if result != ResultFull {
			t2.Errorf("Expected ResultFull but have %d (%s)", result, result)
		}
	})
}
