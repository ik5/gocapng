//go:build linux

package gocapng

// Type of acts that are supported
const (
	// Drop the capabilities settings
	ActDrop Act = iota
	// Add the capabilities settings
	ActAdd
)

// Operation types that can be performed (including bitwise)
const (
	TypeEffective   Type = 1
	TypePermitted   Type = 2
	TypeInheritable Type = 4
	TypeBoundingSet Type = 8
	TypeAmbient     Type = 16
)

// type of selects
const (
	SelectCaps    Select = 16
	SelectBounds  Select = 32
	SelectBoth    Select = 48
	SelectAmbient Select = 64
	SelectAll     Select = 112
)

// Result status
const (
	ResultFail Result = iota - 1
	ResultNone
	ResultPartial
	ResultFull
)

// Print types
const (
	PrintStdOut Print = iota
	PrintBuffer
)

// Supported flags
const (
	// Simply change uid and retain specified capabilities and that's all.
	FlagsNoFlag Flags = 0
	// After changing id, remove any supplement groups that may still be in effect from the old uid.
	FlagsDropSuppGrp Flags = 1
	// Clear the bounding set regardless to the internal representation already setup prior to changing the uid/gid.
	FlagsClearBounding Flags = 2
	// After changing id, initialize any supplement groups that may come with the new account.
	// If given with CAPNG_DROP_SUPP_GRP it will have no effect.
	FlagsInitSuppGrp Flags = 4
	// Clear ambient capabilities regardless of the internal representation already setup prior to changing the uid/gid.
	FlagsClearAmbient Flags = 8
)

// UnsetRootID for namespace root id
const UnsetRootID int = -1

// SupportsAmbient to support new (libcap-ng) cap and not libcap
const SupportsAmbient int = 1
