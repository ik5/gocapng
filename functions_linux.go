//go:build linux && cgo

package gocapng

// #include <stdlib.h>
// #include <cap-ng.h>
// #include <stdarg.h>
// #cgo LDFLAGS: -lcap-ng
//
// int capng_updatev_wrapper(capng_act_t action, capng_type_t type, int *capability) {
//	return capng_update(action, type, capability[0]);
//}
import "C"
import (
	"os"
	"unsafe"
)

// CapNG implement the binding of libcap-ng by initialize the .so binding.
// When done using you must use the Close function.
type CapNG struct{}

// Init initialize the pointer for all supported functions
func Init() *CapNG {
	return &CapNG{}
}

// Clear clears chosen capabilities set
//
// Clear sets to 0 all bits in the selected posix capabilities set.
// The options are SelectCaps for the traditional capabilities, SelectBounds
// for the bounding set, SelectBoth if clearing both is desired, SelectAmbient
// if only operating on the ambient capabilities, or SelectAll if clearing all
// is desired.
func (cp CapNG) Clear(set Select) {
	C.capng_clear(C.capng_select_t(set))
}

// Fill chosen capabilities set
//
// Fill sets all bits to a 1 in the selected POSIX capabilities set. The
// options are SelectCaps for the traditional capabilities, SelectBounds for
// the bounding set, SelectBoth if filling both is desired, SelectAmbient if
// only operating on the ambient capabilities, or SelectAll if clearing all is
// desired.
func (cp CapNG) Fill(set Select) {
	C.capng_fill(C.capng_select_t(set))
}

// SetPID  set working pid.
//
// sets the working pid for capabilities operations. This is useful if you want
// to get the capabilities of a different process.
func (cp CapNG) SetPID(pid int) {
	C.capng_setpid(C.int(pid))
}

// GetCapsProcess get the capabilities from a process.
//
// GetCapsProcess will get the capabilities and bounding set of the pid stored
// inside libcap-ng's state table. The default is the pid of the running process.
// This can be changed by using the capng_setpid function.
func (cp CapNG) GetCapsProcess() bool {
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
func (cp CapNG) Update(action Act, t Type, capability Capability) bool {
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
func (cp CapNG) Updatev(action Act, t Type, capability ...Capability) bool {
	if len(capability) == 0 {
		return false
	}
	var caps []C.int

	for _, cap := range capability {
		caps = append(caps, C.int(cap))
	}

	caps = append(caps, C.int(-1))

	result := C.capng_updatev_wrapper(
		C.capng_act_t(action),
		C.capng_type_t(t),
		(*C.int)(unsafe.Pointer(&caps[0])),
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
func (cp CapNG) Apply(set Select) error {
	result := C.capng_apply(C.capng_select_t(set))

	switch result {
	case -1:
		return ErrNotInitialized
	case -2:
		return ErrSelectBoundsFailureDropBoundingSetCapability
	case -3:
		return ErrSelectBoundsAndFailureToReReadBoundingSet
	case -4:
		return ErrSelectBoundsCAPSetPCap
	case -5:
		return ErrSelectCapsCapsetSyscall
	case -6:
		return ErrSelectAmbientAndProcessClearing
	case -7:
		return ErrSelectAmbientProcessCapabilitiesClearing
	case -8:
		return ErrSelectAmbientProcessCapabilitiesSetting
	}

	return nil
}

// Lock locks the current process capabilities settings
//
// lock will take steps to prevent children of the current process to regain
// full privileges if the uid is 0. This should be called while possessing the
// CAPSetPCap capability in the kernel. This function will do the following if
// permitted by the kernel: Set the NOROOT option on for  PR_SET_SECUREBITS, set
// the NOROOT_LOCKED option to on for PR_SET_SECUREBITS, set the PR_NO_SETUID_FIXUP
// option on for PR_SET_SECUREBITS, and set the PR_NO_SETUID_FIXUP_LOCKED option
// on for PR_SET_SECUREBITS.
func (cp CapNG) Lock() bool {
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
func (cp CapNG) ChangeID(uid, gid int, flag Flags) error {
	result := C.capng_change_id(C.int(uid), C.int(gid), C.capng_flags_t(flag))
	switch result {
	case -1:
		return ErrCAPNGNotInittedProperly
	case -2:
		return ErrFailureRequestingCapabilitiesUidChange
	case -3:
		return ErrApplyingIntermediateCapabilitiesFailed
	case -4:
		return ErrChangingGIDFailed
	case -5:
		return ErrDroppingSupplementalGroupsFailed
	case -6:
		return ErrChangingUIDFailed
	case -7:
		return ErrDroppingAbilityRetainUIDChangeFailed
	case -8:
		return ErrClearingBoundingSet
	case -9:
		return ErrDroppingCAPSETPCAP
	case -10:
		return ErrInitializedSupplementalGroups
	}
	return nil
}

// GetRootID - get namespace root id
// capng_get_rootid gets the rootid for capabilities operations. This is only
// applicable for file system operations.
//
// If the file is in the init namespace or the kernel does not support V3 file
// system capabilities, it returns UnsetRootID. Otherwise it return an integer
// for the namespace root id.
func (cp CapNG) GetRootID() int {
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
func (cp CapNG) SetRootID(rootID int) bool {
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
func (cp CapNG) GetCapsFD(fd os.File) bool {
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
func (cp CapNG) ApplyCapsFD(fd os.File) error {
	intFD := int(fd.Fd())
	result := C.capng_apply_caps_fd(C.int(intFD))

	switch result {
	case -1:
		return ErrFDIsNotRegularFile
	case -2:
		return ErrNonRootNamespaceIDUsedForRootID
	}
	return nil
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
func (cp CapNG) HaveCapabilities(set Select) Result {
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
func (cp CapNG) HavePermittedCapabilities() Result {
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
func (cp CapNG) HaveCapability(which Type, capability Capability) bool {
	result := C.capng_have_capability(
		C.capng_type_t(which),
		C.uint(capability),
	)
	return result == 1
}

// PrintCapsNumeric print numeric values for capabilities set
//
// PrintCapsNumeric will create a numeric representation of the internal
// capabilities. The representation can be sent to either stdout or a buffer by
// passing PrintStdOut or PrintBuffer respectively for the where parameter.
//
// The set parameter controls what is included in the representation. The legal
// options are SelectCaps for the traditional capabilities, SelectBounds for the
// bounding set, SelectBoth if printing both is desired, SelectAmbient if only
// printing the ambient capabilities, or SelectAll if printing all is desired.
//
// If PrintBuffer was selected for where, this will be the text buffer and NULL
// on failure. If PrintStdOut was selected then this value will be NULL no matter
// what.
func (cp CapNG) PrintCapsNumberic(where Print, set Select) string {
	result := C.capng_print_caps_numeric(
		C.capng_print_t(where),
		C.capng_select_t(set),
	)

	if result == nil {
		return ""
	}

	str := C.GoString(result)

	if where == PrintBuffer {
		// Based on man file capng_print_caps_numeric(3):
		//
		// If the option was for a buffer, this function will malloc a buffer that
		// the caller must free.
		C.free(unsafe.Pointer(result))
	}

	return str
}

// PrintCapsText print names of values for capabilities set
//
// PrintCapsText will create a text string representation of the internal
// capability set specified. The representation can be sent to either stdout or
// a buffer by passing PrintStdOut or PrintBuffer respectively for the where
// parameter.
//
// The legal values for the which parameter is Typeeffective, TypePermitted,
// TypeInheritable, TypeBoundingSet, or TypeAmbient.
//
// If PrintBuffer was selected for where, this will be the string buffer and
// empty string on failure. If PrintStdOut was selected then this value will be
// empty string no matter what.
func (cp CapNG) PrintCapsText(where Print, which Type) string {
	result := C.capng_print_caps_text(
		C.capng_print_t(where),
		C.capng_type_t(which),
	)

	if result == nil {
		return ""
	}

	str := C.GoString(result)

	if where == PrintBuffer {
		// Based on man file capng_print_caps_numeric(3):
		//
		// If the option was for a buffer, this function will malloc a buffer that
		// the caller must free.
		C.free(unsafe.Pointer(result))
	}

	return str
}

// NameToCapability  convert capability text to integer
//
// NameToCapability will take the string being passed and look it up to see what
// its integer value would be. The string being input is the same name as the
// define in linux/capabiliy.h with the CAP_ prefix removed. The string case does
// not matter. The integer that is output is the same as the define would be from
// linux/capabiliy.h. This is useful for taking string input and converting to
// something that can be used with Update.
//
// This returns a Capability and nil error, or an error not found and 0 on
// capability.
func (cp CapNG) NameToCapability(name string) (Capability, error) {
	namePtr := C.CString(name)
	defer C.free(unsafe.Pointer(namePtr))

	result := C.capng_name_to_capability(namePtr)
	if result == -1 {
		return 0, ErrCapabilityNotFound
	}
	return Capability(result), nil
}

// CapabilityToName convert capability integer to text
//
// CapabilityToName will take the integer being passed and look it up to see
// what its text string representation would be. The integer being input must be
// in the valid range defined in linux/capabiliy.h. The string that is output is
// the same as the define text from linux/capabiliy.h with the CAP_ prefix
// removed and lower case. This is useful for taking integer representation and
// converting it to something more user friendly for display.
func (cp CapNG) CapabilityToName(capability Capability) string {
	result := C.capng_capability_to_name(C.uint(capability))
	if result == nil {
		return ""
	}
	name := C.GoString(result)
	return name
}
