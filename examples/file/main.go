package main

import (
	"fmt"
	"os"

	"github.com/ik5/gocapng"
)

func main() {
	cap := gocapng.Init() // initialize libcap-ng

	// try to get the current process capabilities (type of initiation)
	if !cap.GetCapsProcess() {
		fmt.Println("Unable to get process capabilities")
		os.Exit(1)
	}

	// Any thread and sub process will inherit settings
	applyTo := gocapng.TypeInheritable
	if !cap.Update(
		gocapng.ActAdd,
		applyTo,
		gocapng.CAPSetPCap, // make sure that we can set capabilities
	) {
		fmt.Println("Unable to set CAPSetPCap")
		os.Exit(1)
	}

	if !cap.Update(
		gocapng.ActAdd,
		applyTo,
		gocapng.CAPSetFCap, // make sure we can set file capabilities
	) {
		fmt.Println("Unable to set CAPSetFCap")
		os.Exit(1)
	}

	err := cap.Apply(gocapng.SelectAmbient) // apply the given request.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to apply capability: %s\n", err)
		os.Exit(2)
	}

	f, err := os.OpenFile("/var/run/file-example.pid", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create pid file: %s\n", err)
		os.Exit(3)
	}
	defer f.Close()

	pid := os.Getpid()
	fmt.Println("my pid:", pid)
	n, err := fmt.Fprintf(f, "%d", pid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Written", n, "bytes")
	f.Sync()
}
