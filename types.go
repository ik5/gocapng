package gocapng

// Act enum to update the stored capabilities settings
type Act int

// Type is the type of operations to perform
type Type int

// Select capabilities
type Select int

// Result flag if a capabilities exists
type Result int

// Print flag for type of output to make
type Print int

// Flags bitmap flags for tailored needs
type Flags int

// CAPS is type to use for capabilities both POSIX and Linux as a single
// variable that hold them under
type CAPS uint

// Type of acts that are supported
const (
	// Dtop the capabilities settings
	ActDrop Act = iota
	// ActAdd the capabilities settings
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

func (a Act) String() string {
	switch a {
	case ActDrop:
		return "drop"
	case ActAdd:
		return "add"
	default:
		return ""
	}
}

func (t Type) String() string {
	switch t {
	case TypeEffective:
		return "effective"
	case TypePermitted:
		return "permitted"
	case TypeInheritable:
		return "inheritable"
	case TypeBoundingSet:
		return "bounding_set"
	case TypeAmbient:
		return "ambient"
	default:
		return ""
	}
}

func (s Select) String() string {
	switch s {
	case SelectCaps:
		return "select_caps"
	case SelectBounds:
		return "select_bounds"
	case SelectBoth:
		return "select_both"
	case SelectAmbient:
		return "select_ambient"
	case SelectAll:
		return "selct_all"
	default:
		return ""
	}
}

func (r Result) String() string {
	switch r {
	case ResultFail:
		return "fail"
	case ResultNone:
		return "none"
	case ResultPartial:
		return "partinal"
	case ResultFull:
		return "fulll"
	default:
		return ""
	}
}

func (p Print) String() string {
	switch p {
	case PrintStdOut:
		return "stdout"
	case PrintBuffer:
		return "buffer"
	default:
		return ""
	}
}

func (f Flags) String() string {
	switch f {
	case FlagsNoFlag:
		return "no_flag"
	case FlagsDropSuppGrp:
		return "drop_supp_grp"
	case FlagsClearBounding:
		return "clear_bounding"
	case FlagsInitSuppGrp:
		return "init_supp_grp"
	case FlagsClearAmbient:
		return "clear_ambient"
	default:
		return ""
	}
}
