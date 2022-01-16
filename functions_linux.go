//go:build linux && cgo

package gocapng

// #include <stdlib.h>
// #include <cap-ng.h>
// #cgo LDFLAGS: -lcap-ng
import "C"

// CapsNG implement the binding of libcap-ng by initialize the .so binding.
// When done using you must use the Close function.
type CapsNG struct {
}

// InitFunctions initialize the pointer for all supported functions
func InitFunctions() (*CapsNG, error) {
	return &CapsNG{}, nil
}

// Clear clears chosen capabilities set
//
// Clear sets to 0 all bits in the selected posix capabilities set.
// The options are SelectCaps for the traditional capabilities, SelectBounds
// for the bounding set, SelectBoth if clearing both is desired, SelectAmbient
// if only operating on the ambient capabilities, or SelectAll if clearing all
// is desired.
func (cp CapsNG) Clear(set Select) {
	C.capng_clear(C.capng_select_t(set))
}

// Fill chosen capabilities set
//
// Fill sets all bits to a 1 in the selected posix capabilities set. The
// options are SelectCaps for the traditional capabilities, SelectBounds for
// the bounding set, SelectBoth if filling both is desired, SelectAmbient if
// only operating on the ambient capabilities, or SelectAll if clearing all is
// desired.
func (cp CapsNG) Fill(set Select) {
	C.capng_fill(C.capng_select_t(set))
}

// SetPID  set working pid.
//
// sets the working pid for capabilities operations. This is useful if you want
// to get the capabilities of a different process.
func (cp CapsNG) SetPID(pid int) {
	C.capng_setpid(C.int(pid))
}

// GetCapsProcess get the capabilities from a process.
//
// GetCapsProcess will get the capabilities and bounding set of the pid stored
// inside libcap-ng's state table. The default is the pid of the running process.
// This can be changed by using the capng_setpid function.
func (cp CapsNG) GetCapsProcess() bool {
	result := C.capng_get_caps_process()
	return int(result) == 0
}

// Update update the stored capabilities settings.
//
// capng_update  will update the internal posix capabilities settings based on
// the options passed to it. The action should be either ActDrop to set the
// capability bit to 0, or ActAdd to set the capability bit to 1. The operation
// is performed on the capability set specified in the type parameter. The
// values are: TypeEffective, Typepermitted, TypeInheritable, TypeBoundingSet,
// or TypeAmbient. The values may be or'ed together to perform the same operation
// on multiple sets. The last parameter, capability, is the capability define as
// given in linux/capability.h (translated into Golang by this package.
//
// This returns 0 on success and -1 on failure.
func (cp CapsNG) Update(action Act, t Type, capability CAPS) bool {
	result := C.capng_update(
		C.capng_act_t(action), C.capng_type_t(t), C.uint(capability),
	)

	return int(result) == 0
}
