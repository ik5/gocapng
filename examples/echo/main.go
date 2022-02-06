package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ik5/gocapng"
)

const (
	echoPort = "7"
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

func echo() {

}

func listenServer(quit chan bool) {
	fmt.Println("Initializing Echo TCP Server")
	tcpListen, err := net.Listen("tcp", ":"+echoPort)
	if err != nil {
		quit <- true
		fmt.Fprintf(os.Stderr, "Unable to bind to on ':%s': %s", echoPort, err)
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

			fmt.Fprintf(conn, "%s", conn)
		}
	}

}

func main() {
	cap := gocapng.Init()        // initialize libcap-ng
	cap.Clear(gocapng.SelectAll) // clear memory, and set everything to "nothing"

	if !cap.Updatev(
		gocapng.ActAdd,
		gocapng.TypeEffective|gocapng.TypePermitted|gocapng.TypeBoundingSet,
		gocapng.CAPSetPCap,
		gocapng.CAPNetBindService,
		// gocapng.CAPSetFCap,
	) {
		fmt.Println("Unable to request capability for binding low port number.")
		os.Exit(-1)
	}

	err := cap.Apply(gocapng.SelectBoth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to apply capability: %s\n", err)
		os.Exit(-2)
	}

	quit := make(chan bool)
	go handleSignals(quit)
	go listenServer(quit)

	<-quit
}
