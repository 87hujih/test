package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	for i := 0; i < 5; i++ {
		conn, _ := net.DialTCP("tcp4", nil, addr)
		conn.Write([]byte("发送了消息" + strconv.Itoa(i)))
		a := make([]byte, 1024)
		read, _ := conn.Read(a)
		fmt.Println("服务器返回了消息", string(a[:read]))
		conn.Close()
	}
}
