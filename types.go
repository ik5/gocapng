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

// UserCapData holds libcap user data
type UserCapData struct {
	Effective   uint32
	Permitted   uint32
	inheritable uint32
}

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
