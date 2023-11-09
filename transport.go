package main

import (
	"encoding/json"
	"fmt"
	N "github.com/NullpointerW/anicat/net"
	"github.com/NullpointerW/anicat/net/cmd"

	// "log"
	"net"
)

func Send(dialAddress string, cmd cmd.Cmd) (string, error) {
	signal := make(chan struct{})
	go waitProgress(signal)
	c, err := net.Dial("tcp", dialAddress)
	if err != nil {
		return "", err
	}
	nc := &N.Conn{
		TcpConn: c,
		Max:     1024 * 1024,
	}
	b, _ := json.Marshal(cmd)
	err = nc.Write(string(b))
	if err != nil {
		return "", err
	}
	signal <- struct{}{}
	read, err := nc.Read()
	if err != nil {
		return "", err
	}
	signal <- struct{}{}
	fmt.Print(clearLine)
	fmt.Print(cursorVisible)
	return string(read), nil
}
