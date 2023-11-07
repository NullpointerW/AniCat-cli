package main

import (
	"encoding/json"
	N "github.com/NullpointerW/anicat/net"
	// "log"
	"net"
)

type cTyp int

const (
	Add cTyp = iota
	AddFeed
	Remove
	Ls
	LsItems
	Status
	Stop
)

type Cmd struct {
	Cmd cTyp            `json:"cmd"`
	Arg string          `json:"arg"`
	Raw json.RawMessage `json:"raw"`
}

func Send(dialAddress string, cmd Cmd) (string, error) {
	signal := make(chan struct{})
	go waitProgress(signal)
	c, err := net.Dial("tcp", dialAddress)
	if err != nil {
		return "", err
	}
	nc := &N.Conn{
		TcpConn: c,
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
	return read, nil
}
