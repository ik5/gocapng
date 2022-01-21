//go:build linux && cgo

package gocapng

// #include <stdlib.h>
// #include <cap-ng.h>
// #include <stdarg.h>
// #cgo LDFLAGS: -lcap-ng
//
// int capng_updatev_wrapper(capng_act_t action, capng_type_t type, unsigned int *capability) {
//	return capng_update(action, type, capability[0]);
//}
import "C"
import (
	"errors"
	"os"
	"unsafe"
)

// CapsNG implement the binding of libcap-ng by initialize the .so binding.
// When done using you must use the Close function.
type CapsNG struct{}

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
// given in linux/capability.h (translated into Golang by this package).
//
// This returns true on success and false on failure.
func (cp CapsNG) Update(action Act, t Type, capability CAPS) bool {
	result := C.capng_update(
		C.capng_act_t(action),
		C.capng_type_t(t),
		C.uint(capability),
	)

	return int(result) == 0
}

// Updatev  update the stored capabilities settings
// updatev will update the internal posix capabilities settings based on the
// options passed to it. The action should be either ActDrop to set the
// capability bit to 0, or ActAdd to set the capability bit to 1. The operation
// is performed on the capability set specified in the type parameter.
// The values are: TypeEffective, TypePermitted, TypeInheritable, TypeBoundingSet,
// or TypeAmbient.
// The values may be or'ed together to perform the same operation on multiple
// sets. The last parameter, capability, is the capability define as given in
// linux/capability.h (translated into Golang by this package).
//
// This function differs from update in that you may pass a list of
// capabilities.
//
// This returns true on success and false on failure.
func (cp CapsNG) Updatev(action Act, t Type, capability ...CAPS) bool {
	if len(capability) == 0 {
		return false
	}
	var caps []C.int

	for _, cap := range capability {
		caps = append(caps, C.int(cap))
	}

	caps = append(caps, -1)

	result := C.capng_updatev_wrapper(
		C.capng_act_t(action),
		C.capng_type_t(t),
		(*C.uint)(unsafe.Pointer(&caps)),
	)

	return int(result) == 0
}

// Apply the stored capabilities settings.
// capng_apply will transfer the specified internal posix capabilities settings
// to the kernel. The options are SelectCaps for the traditional capabilities,
// SelectBounds for the bounding set, SelectBoth if transferring both is desired,
// SelectAmbient if only operating on the ambient capabilities, or SelectAll if
// applying all is desired.
//
//
func (cp CapsNG) Apply(set Select) error {
	var err error
	result := C.capng_apply(C.capng_select_t(set))

	switch result {
	case -1:
		err = errors.New("not initialized")
	case -2:
		err = errors.New("SelectBounds and failure to drop a bounding set capability")
	case -3:
		err = errors.New("SelectBounds and failure to re-read bounding set")
	case -4:
		err = errors.New("SelectBounds and process does not have CAPSetPCap")
	case -5:
		err = errors.New("SelectCaps and failure in capset syscall")
	case -6:
		err = errors.New("SelectAmbient and process has no capabilities and failed clearing ambient capabilities")
	case -7:
		err = errors.New("SelectAmbient and process has capabilities and failed clearing ambient capabilities")
	case -8:
		err = errors.New("SelectAmbient and process has capabilities and failed setting an ambient capability")
	}

	return err
}

// Lock locks the current process capabilities settings
// lock will take steps to prevent children of the current process to regain
// full privileges if the uid is 0. This should be called while possessing the
// CAPSetPCap capability in the kernel. This function will do the following if
// permitted by the kernel: Set the NOROOT option on for  PR_SET_SECUREBITS, set
// the NOROOT_LOCKED option to on for PR_SET_SECUREBITS, set the PR_NO_SETUID_FIXUP
// option on for PR_SET_SECUREBITS, and set the PR_NO_SETUID_FIXUP_LOCKED option
// on for PR_SET_SECUREBITS.
func (cp CapsNG) Lock() bool {
	result := C.capng_lock()
	return result == 0
}

// ChangeID  changes the credentials retaining capabilities
//  This function will change uid and gid to the ones given while retaining the
// capabilities previously specified in capng_update. It is also possible to
// specify -1 for either the uid or gid in which case the function will not
// change the uid or gid and leave it "as is". This is useful if you just want
// the flag options to be applied (assuming the option doesn't require more
// privileges that you currently have).
//
// It is not necessary and perhaps better if Apply has not been called prior to
// this function so that all necessary privileges are still intact. The
// caller may be required to have CAPSetPCap capability still active before
// calling this function or capabilities cannot be changed.
//
// This function also takes a flag parameter that helps to tailor the exact
// actions performed by the function to secure the environment.
// The option may be or'ed together. The legal values are:
//
//    FlagNoFlag
//        Simply change uid and retain specified capabilities and that's all.
//
//    FlagDropSuppGrp
//        After changing id, remove any supplement groups that may still be in
//        effect from the old uid.
//
//    FlagInitSuppGrp
//        After changing id, initialize any supplement groups that may come with
//        the new account. If given with CAPNG_DROP_SUPP_GRP it will have no effect.
//
//    FlagClearBounding
//        Clear the bounding set regardless to the internal representation
//        already setup prior to changing the uid/gid.
//
//    FlagClearAmbient
//        Clear ambient capabilities regardless of the internal representation
//        already setup prior to changing the uid/gid.
func (cp CapsNG) ChangeID(uid, gid int, flag Flags) error {
	var err error

	result := C.capng_change_id(C.int(uid), C.int(gid), C.capng_flags_t(flag))
	switch result {
	case -1:
		err = errors.New("means capng has not been initted properly")
	case -2:
		err = errors.New("means a failure requesting to keep capabilities across the uid change")
	case -3:
		err = errors.New("means that applying the intermediate capabilities failed")
	case -4:
		err = errors.New("means changing gid failed")
	case -5:
		err = errors.New("means dropping supplemental groups failed")
	case -6:
		err = errors.New("means changing the uid failed")
	case -7:
		err = errors.New("means dropping the ability to retain caps across a uid change failed")
	case -8:
		err = errors.New("means clearing the bounding set failed")
	case -9:
		err = errors.New("means dropping CAP_SETPCAP or ambient capabilities failed")
	case -10:
		err = errors.New("means initializing supplemental groups failed")
	}
	return err
}

// GetRootID - get namespace root id
// capng_get_rootid gets the rootid for capabilities operations. This is only
// applicable for file system operations.
//
// If the file is in the init namespace or the kernel does not support V3 file
// system capabilities, it returns UnsetRootID. Otherwise it return an integer
// for the namespace root id.
func (cp CapsNG) GetRootID() int {
	result := C.capng_get_rootid()
	return int(result)
}

// SetRootID set namespace root id
//SetRootID sets the rootid for capabilities operations. This is only
// applicable for file system operations.
//
// On false there is an internal error or the kernel does not suppor V3
// filesystem capabilities. On false f there is an internal error or the kernel
// does not suppor V3 // filesystem capabilities.
func (cp CapsNG) SetRootID(rootID int) bool {
	result := C.capng_set_rootid(C.int(rootID))
	return result == 0
}

// GetCapsFD Read file based capabilities
//
// This function  will read the file based capabilities stored in extended
// attributes of the file that the descriptor was opened against. The bounding
// set is not included in file based capabilities operations. Note that this
// function will only work if compiled on a kernel that supports file based
// capabilities such as 2.6.26 and later. If the "magic" bit is set, then all
// effect capability bits are set. Otherwise the bits are cleared.
func (cp CapsNG) GetCapsFD(fd os.File) bool {
	intFD := int(fd.Fd())
	result := C.capng_get_caps_fd(C.int(intFD))
	return result == 0
}

// ApplyCapsFD writes the capabilities for a file.
//
// This function will write the file based capabilities to the extended
// attributes of the file that the descriptor was opened against. The bounding
// set is not included in file based capabilities operations. Note that this
// function will only work if compiled on a kernel that supports file based
// capabilities such as 2.6.2 6 and later.
func (cp CapsNG) ApplyCapsFD(fd os.File) error {
	intFD := int(fd.Fd())
	result := C.capng_apply_caps_fd(C.int(intFD))

	var err error
	switch result {
	case -1:
		err = errors.New("fd is not a regular file")
	case -2:
		err = errors.New("non-root namespace id is being used for rootid")
	}
	return err
}

// HaveCapabilities check for capabilities
//
// HaveCapabilities will  check the selected internal capabilities sets to
// see what the status is. The capabilities sets must be previously setup with
// calls to GetCapsProcess, GetCapsFD, or in some other way setup. The options
// are SelectCaps for the traditional capabilities, SelectBounds for the
// bounding set, SelectBoth if checking both are desired, SelectAmbient if only
// checking the ambient capabilities, or SelectAll if testing all sets is desired.
// When capabilities are checked, it will only look at the effective capabilities.
//
// Will not work for a file, use HavePermittedCapabilities instead.
func (cp CapsNG) HaveCapabilities(set Select) Result {
	result := C.capng_have_capabilities(C.capng_select_t(set))
	return Result(result)
}

// HavePermittedCapabilities check for capabilities
//
// HavePermittedCapabilities will check the selected internal capabilities sets to
// see what the status is. The capabilities sets must be previously setup with
// calls to GetCapsProcess, GetCapsFD, or in some other way setup. The options
// are SelectCaps for the traditional capabilities, SelectBounds for the
// bounding set, SelectBoth if checking both are desired, SelectAmbient if only
// checking the ambient capabilities, or SelectAll if testing all sets is desired.
// When capabilities are checked, it will only look at the effective capabilities.
//
// The source of capabilities comes from a file, then you may need to additionally
// check the permitted capabilities. It's for this reason that
// HavePermittedCapabilities was created. It takes no arguments because it
// simply checks the permitted set.
func (cp CapsNG) HavePermittedCapabilities() Result {
	result := C.capng_have_permitted_capabilities()
	return Result(result)
}

// HaveCapability check for specific capability
//
// HaveCapability will check the specified internal capabilities set to see if
// the specified capability is set. The capabilities sets must be previously
// setup with calls to GetCapsProcess, GetCapsFD, or in some other way setup.
// The values for which should be one of: TypeEffective, TypePermitted,
// TypeInheritable, TypeBounding_set, or TypeAmbient.
func (cp CapsNG) HaveCapability(which Type, capability uint) bool {
	result := C.capng_have_capability(
		C.capng_type_t(which),
		C.uint(capability),
	)
	return result == 1
}
