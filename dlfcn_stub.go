//go:build !linux

package gocapng

import "errors"

type dynamicLibrary struct{}

var notSupported = errors.New("not supported")

func newLibrary(name string, flags int) *dynamicLibrary {
	return nil
}

func (dl *dynamicLibrary) Open() error {
	return notSupported
}

func (dl *dynamicLibrary) Close() error {
	return notSupported
}

func (dl *dynamicLibrary) Lookup(symbol string) error {
	return notSupported
}
