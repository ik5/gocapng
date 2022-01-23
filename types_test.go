package gocapng

import "testing"

func TestActString(t *testing.T) {
	toCheck := []struct {
		act      Act
		name     string
		expected string
	}{
		{
			act:      ActDrop,
			name:     "ActDrop",
			expected: "drop",
		},
		{
			act:      ActAdd,
			name:     "ActAdd",
			expected: "add",
		},
		{
			act:      Act(-1),
			name:     "-1",
			expected: "",
		},
	}

	for _, act := range toCheck {
		if act.act.String() != act.expected {
			t.Errorf(
				"'%s expected to be '%s' but have '%s' instead",
				act.name, act.expected, act.act.String(),
			)
		}
	}
}

func TestTypeString(t *testing.T) {
	toCheck := []struct {
		t        Type
		name     string
		expected string
	}{
		{
			t:        TypeEffective,
			name:     "TypeEffective",
			expected: "effective",
		},
		{
			t:        TypePermitted,
			name:     "TypePermitted",
			expected: "permitted",
		},
		{
			t:        TypeInheritable,
			name:     "TypeInheritable",
			expected: "inheritable",
		},
		{
			t:        TypeBoundingSet,
			name:     "TypeBoundingSet",
			expected: "bounding_set",
		},
		{
			t:        TypeAmbient,
			name:     "TypeAmbient",
			expected: "ambient",
		},
		{
			t:        Type(-1),
			name:     "-1",
			expected: "",
		},
	}

	for _, types := range toCheck {
		if types.t.String() != types.expected {
			t.Errorf(
				"'%s expected to be '%s' but have '%s' instead",
				types.name, types.expected, types.t.String(),
			)
		}
	}
}

func TestSelectString(t *testing.T) {
	toCheck := []struct {
		s        Select
		name     string
		expected string
	}{
		{
			s:        SelectCaps,
			name:     "SelectCaps",
			expected: "select_caps",
		},
		{
			s:        SelectBounds,
			name:     "SelectBounds",
			expected: "select_bounds",
		},
		{
			s:        SelectBoth,
			name:     "SelectBoth",
			expected: "select_both",
		},
		{
			s:        SelectAmbient,
			name:     "SelectAmbient",
			expected: "select_ambient",
		},
		{
			s:        SelectAll,
			name:     "SelectAll",
			expected: "select_all",
		},
		{
			s:        Select(-1),
			name:     "-1",
			expected: "",
		},
	}

	for _, sel := range toCheck {
		if sel.s.String() != sel.expected {
			t.Errorf(
				"'%s expected to be '%s' but have '%s' instead",
				sel.name, sel.expected, sel.s.String(),
			)
		}
	}
}

func TestResultString(t *testing.T) {
	toCheck := []struct {
		r        Result
		name     string
		expected string
	}{
		{
			r:        ResultFail,
			name:     "ResultFail",
			expected: "fail",
		},
		{
			r:        ResultNone,
			name:     "ResultNone",
			expected: "none",
		},
		{
			r:        ResultPartial,
			name:     "ResultPartial",
			expected: "partial",
		},
		{
			r:        ResultFull,
			name:     "ResultFull",
			expected: "full",
		},
		{
			r:        Result(-2),
			name:     "-2",
			expected: "",
		},
	}
	for _, r := range toCheck {
		if r.r.String() != r.expected {
			t.Errorf(
				"'%s expected to be '%s' but have '%s' instead",
				r.name, r.expected, r.r.String(),
			)
		}
	}
}

func TestPrintString(t *testing.T) {
	toCheck := []struct {
		p        Print
		name     string
		expected string
	}{
		{
			p:        PrintStdOut,
			name:     "PrintStdOut",
			expected: "stdout",
		},
		{
			p:        PrintBuffer,
			name:     "PrintBuffer",
			expected: "buffer",
		},
		{
			p:        Print(-1),
			name:     "-1",
			expected: "",
		},
	}

	for _, print := range toCheck {
		if print.p.String() != print.expected {
			t.Errorf("'%s' expected '%s' got '%s'",
				print.name, print.expected, print.p.String(),
			)
		}
	}
}

func TestFlagsString(t *testing.T) {
	toCheck := []struct {
		f        Flags
		name     string
		expected string
	}{
		{
			f:        FlagsNoFlag,
			name:     "FlagsNoFlag",
			expected: "no_flag",
		},
		{
			f:        FlagsDropSuppGrp,
			name:     "FlagsDropSuppGrp",
			expected: "drop_supp_grp",
		},
		{
			f:        FlagsClearBounding,
			name:     "FlagsClearBounding",
			expected: "clear_bounding",
		},
		{
			f:        FlagsInitSuppGrp,
			name:     "FlagsInitSuppGrp",
			expected: "init_supp_grp",
		},
		{
			f:        FlagsClearAmbient,
			name:     "FlagsClearAmbient",
			expected: "clear_ambient",
		},
		{
			f:        Flags(-1),
			name:     "-1",
			expected: "",
		},
	}

	for _, f := range toCheck {
		if f.f.String() != f.expected {
			t.Errorf("'%s' expected '%s' got '%s'",
				f.name, f.expected, f.f.String(),
			)
		}
	}
}
