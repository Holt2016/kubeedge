package server

import (
	"github.com/kubeedge/beehive/pkg/common/log"
	"net"
	"testing"
	"time"
)
func TestDnsStart(t *testing.T) {
	inter = "lo"
	go DnsStart()

	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP:net.IPv4zero, Port:0}
	dstAddr := &net.UDPAddr{IP:sip, Port:53}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		t.Errorf("connect failed, err : %v\n", err.Error())
		return
	}
	defer conn.Close()

	//handle connection read
	done := make(chan string)
	go handleUDPRead(conn, done)

	//write raw bytes to UDP server
	testString := []byte{0x98, 0x6b, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x02, 0x77, 0x33, 0x06,
		0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x03, 0x63,
		0x6f, 0x6d, 0x00, 0x00, 0x01, 0x00, 0x01,
	}
	_, err = conn.Write(testString)
	if err != nil {
		t.Errorf("write failed, err : %v\n", err)
	}
	<-done
	time.Sleep(3 * time.Second)
}

func handleUDPRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.LOGGER.Infof("error to read from UDP server: ", err)
		done <- "done"
		return
	}
	log.LOGGER.Info("UDP server response: " + string(buf[:]))
	done <- "done"
}
