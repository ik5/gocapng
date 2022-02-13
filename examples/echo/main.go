package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ik5/gocapng"
)

const (
	echoPort = 7
)

func handleSignals(quit chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(
		sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT,
	)

	for {
		select {
		case sig := <-sigs:
			fmt.Fprintf(os.Stderr, "\nsignal '%s' was pressed\n", sig)
			quit <- true
			return
		}
	}
}

func echoTCP(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read content: %s", err)
			break
		}
		if n > 0 {
			fmt.Fprintf(conn, "%s", buf[:n])
		}
	}

}

func listenTCPServer(quit chan bool) {
	fmt.Println("Initializing Echo TCP Server")
	tcpListen, err := net.Listen("tcp", fmt.Sprintf(":%d", echoPort))
	if err != nil {
		quit <- true
		fmt.Fprintf(os.Stderr, "Unable to bind to on ':%d': %s", echoPort, err)
		return
	}
	defer tcpListen.Close()

	for {
		select {
		case <-quit:
			return
		default:
			conn, err := tcpListen.Accept()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to accept: %s", err)
				continue
			}

			go echoTCP(conn)
		}
	}
}

func listenUDPServer(quit chan bool) {
	fmt.Println("Initializing Echo UDP Server")
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: echoPort})
	if err != nil {
		quit <- true
		fmt.Fprintf(os.Stderr, "Unable to bind to on ':%d': %s", echoPort, err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		select {
		case <-quit:
			return
		default:
			for {
				n, addr, err := conn.ReadFromUDP(buf[:])
				if errors.Is(err, io.EOF) {
					break
				}

				if err != nil {
					fmt.Fprintf(os.Stderr, "unable to read UDP: %s", err)
					break
				}

				if n > 0 {
					conn.WriteToUDP(buf[:n], addr)
				}
			}
		}
	}
}

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

	if !cap.Update(
		gocapng.ActAdd,
		applyTo,
		gocapng.CAPNetBindService, // allow us to bind < 1024 port number in TCP
	) {
		fmt.Println("Unable to request capability for binding low port number.")
		os.Exit(1)
	}

	if !cap.Update(
		gocapng.ActAdd,
		applyTo,
		gocapng.CAPNetRaw, // allow us to bind < 1024 port number in UDP
	) {
		fmt.Println("Unable to request capability for binding low port number.")
		os.Exit(1)
	}

	err := cap.Apply(gocapng.SelectAmbient) // apply the given request.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to apply capability: %s\n", err)
		os.Exit(2)
	}

	quit := make(chan bool)
	go handleSignals(quit)
	go listenTCPServer(quit)
	go listenUDPServer(quit)

	<-quit
}
