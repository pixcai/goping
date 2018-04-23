package main

import (
	"fmt"
	"net"
	"time"
)

var (
	packet *Packet = NewPacket()
	recv   []byte  = make([]byte, 1024)
)

func main() {
	laddr := net.IPAddr{IP: net.ParseIP("127.0.0.1")}
	raddr := net.IPAddr{IP: net.ParseIP("192.168.0.86")}

	conn, err := net.DialIP("ip:icmp", &laddr, &raddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	conn.Write(packet.Bytes())
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	conn.Read(recv)
}
