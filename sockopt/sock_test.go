package sockopt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"syscall"
	"testing"
	"time"
)

func TestTCPRepair(t *testing.T) {
	go serveRandom()
	// server should be good after a sec :/
	time.Sleep(time.Second * 1)
	client := &http.Client{}
	resp, err := client.Get("http://localhost:6666/random")
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("Resp: %s\n", body)
	resp.Body.Close()
	client.CloseIdleConnections()
	//let some time for the close event
	time.Sleep(time.Second * 1)
}

func random(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello !\n")
}

func setTCPRepair(fd, repair int) {
	err := syscall.SetsockoptInt(fd, syscall.SOL_TCP, 19, repair)
	if err != nil {
		log.Fatalf("Couldn't set repair: %s\n", err)
	} else {
		log.Printf("TCP repair %v", repair)
	}
}

func getTCPRepairWindow(fd int) {
	w, e := GetsockoptTCPRepairWindow(fd, syscall.SOL_TCP, 29)
	if e != nil {
		log.Fatalf("Can't get TCPRepairWindow: %s\n", e)
	} else {
		log.Printf("%+v\n", w)
	}
}

func ConnStateEvent(c net.Conn, e http.ConnState) {
	if e != http.StateNew {
		return
	}

	tcpconn, _ := c.(*net.TCPConn)
	file, _ := tcpconn.File()
	fd := int(file.Fd())

	setTCPRepair(fd, 1)
	getTCPRepairWindow(fd)
	setTCPRepair(fd, 0)
}

func serveRandom() {
	server := http.Server{
		Addr:      "localhost:6666",
		ConnState: ConnStateEvent,
	}

	http.HandleFunc("/random", random)
	log.Fatal(server.ListenAndServe())
}
