package net

import (
	"fmt"
	"net"
)

const (
	protocolTCP   = "tcp"
	tcpServerAddr = "127.0.0.1:8000"
)

func tcpServer(c chan<- struct{}) {
	// 服务端初始化
	fmt.Println("server starting...")
	tcpAddr, err := net.ResolveTCPAddr(protocolTCP, tcpServerAddr)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP(protocolTCP, tcpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("server: listening...")
	c <- struct{}{}

	// 监听 fd, 并获取最新的连接
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	// 读取数据
	fmt.Println("server: read message from client")
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("server: received message ", n, string(buf))

	// 写 response
	n, err = conn.Write([]byte("I received your message"))
	if err != nil {
		panic(err)
	}
	fmt.Println("server: send response to client, ", n)
}

func tcpClient(c <-chan struct{}) {
	fmt.Println("client starting...")
	tcpAddr, err := net.ResolveTCPAddr(protocolTCP, tcpServerAddr)
	if err != nil {
		panic(err)
	}

	select {
	case <-c:
		// 与服务器建立连接
		conn, err := net.DialTCP(protocolTCP, nil, tcpAddr)
		if err != nil {
			panic(err)
		}

		// 写消息
		fmt.Println("client: write message")
		n, err := conn.Write([]byte("hello, girl, you done good!"))
		if err != nil {
			panic(err)
		}
		fmt.Println("client: send message to server ", n)

		// 收到 server 的 response
		buf := make([]byte, 512)
		n, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("client: I received response from server, ", string(buf))
	}
}

func TcpTest() {
	c := make(chan struct{}, 1)
	go func() {
		tcpServer(c)
	}()

	tcpClient(c)
	select {}
}
