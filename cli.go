package main

import (
	"bufio"
	"fmt"
	// "log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	N "github.com/NullpointerW/anicat/net"
	"github.com/NullpointerW/anicat/net/cmd"
)

func main() {
	Execute()
	return
	signal := make(chan struct{})
	go waitProgress(signal)

	r := bufio.NewReader(os.Stdin)
	defer func() {
		r.ReadString('\n')
	}()
	dialadress := host + ":" + strconv.Itoa(port)
	c, err := net.Dial("tcp", dialadress)
	if err != nil {
		fmt.Println(cmd.Red, err, cmd.Reset)
		exit(r)
	}
	s := bufio.NewScanner(c)
	s.Split(N.ScanCRLF)
	buf := make([]byte, 0, 64*1024)
	s.Buffer(buf, 1024*1024)
	var f bool
	for s.Scan() {
		if f {
			signal <- struct{}{}
			fmt.Print(clearLine)
			fmt.Print(cursorVisible)
		}
		f = true
		fmt.Println(s.Text())
		if s.Text() == "exited." {
			return
		}
		var (
			err error
			l   string
		)
		for {
			fmt.Print(cmd.Cyan, cmd.Cursor, cmd.Reset)
			l, err = r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			l = string(N.DropCR([]byte(l[:len(l)-1])))
			if l != "cls" && l != "clear" {
				break
			}
			var cmd *exec.Cmd
			if runtime.GOOS == "windows" {
				cmd = exec.Command("cmd", "/c", "cls")
			} else {
				cmd = exec.Command("clear")
			}
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		c.Write([]byte(l + N.CRLF))
		signal <- struct{}{}
	}
	if f {
		signal <- struct{}{}
	}
	fmt.Println(cmd.Red, s.Err(), cmd.Reset)
}

func exit(r *bufio.Reader) {
	r.ReadString('\n')
	os.Exit(1)
}
