package main

import (
	"fmt"
	"log"
	"master"
	"net"
	"os"
)

func onAccept(conn net.Conn) {
	buf := make([]byte, 8192)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read over", err)
			break
		}

		conn.Write(buf[0:n])
	}
}

func onClose(conn net.Conn) {
	log.Println("---client onClose---")
}

func main() {
	var logFile = "./log.txt"

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	log.SetOutput(f)
	//log.SetOutput(io.MultiWriter(os.Stderr, f))

	master.OnClose(onClose)
	master.OnAccept(onAccept)

	if len(os.Args) > 1 && os.Args[1] == "alone" {
		addrs := make([]string, 1)
		if len(os.Args) > 2 {
			addrs = append(addrs, os.Args[2])
		} else {
			addrs = append(addrs, "127.0.0.1:8880")
		}

		fmt.Printf("listen:")
		for _, addr := range addrs {
			fmt.Printf(" %s", addr)
		}
		fmt.Println()

		master.NetStart(addrs)
	} else {
		// daemon mode in master framework
		master.NetStart(nil)
	}
}
