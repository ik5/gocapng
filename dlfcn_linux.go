//go:build linux && cgo

package gocapng

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"
import (
	"errors"
	"unsafe"
)

const (
	rtldLAZY     = C.RTLD_LAZY
	rtldNOW      = C.RTLD_NOW
	rtldGLOBAL   = C.RTLD_GLOBAL
	rtldLOCAL    = C.RTLD_LOCAL
	rtldNODELETE = C.RTLD_NODELETE
	rtldNOLOAD   = C.RTLD_NOLOAD
	rtldDEEPBIND = C.RTLD_DEEPBIND
)

type dynamicLibrary struct {
	Name   string
	Flags  int
	handle unsafe.Pointer
}

func newLibrary(name string, flags int) *dynamicLibrary {
	return &dynamicLibrary{
		Name:   name,
		Flags:  flags,
		handle: nil,
	}
}

func (dl *dynamicLibrary) Open() error {
	handle := C.dlopen(C.CString(dl.Name), C.int(dl.Flags))
	if handle == C.NULL {
		return errors.New(dl.Error())
	}
	dl.handle = handle
	return nil
}

func (dl *dynamicLibrary) Close() error {
	err := C.dlclose(dl.handle)
	if err != 0 {
		return errors.New(dl.Error())
	}
	return nil
}

func (dl *dynamicLibrary) Lookup(symbol string) error {
	C.dlerror() // Clear out any previous errors
	C.dlsym(dl.handle, C.CString(symbol))
	err := C.dlerror()
	if unsafe.Pointer(err) == C.NULL {
		return nil
	}
	return errors.New(C.GoString(err))
}

func (dl *dynamicLibrary) Error() string {
	return C.GoString(C.dlerror())
}
