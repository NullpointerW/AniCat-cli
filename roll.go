package main

import (
	"fmt"
	"time"
)

const (
	clearLine       = "\033[K\r"
	cursorVisible   = "\033[?25h"
	cursorInvisible = "\033[?25l"
)

func display(f string, sec float64, fin chan struct{}) bool {
	defer time.Sleep(60 * time.Millisecond)
	fmt.Printf(f, sec)
	select {
	case <-fin:
		return true
	default:
		return false
	}
}
func waitProgress(c chan struct{}, stop chan struct{}) {
Wait:
	<-c
	st := time.Now()
	for {
		select {
		case <-stop:
			fmt.Print(clearLine)
			fmt.Print(cursorVisible)
			return
		default:
		}
		fmt.Print(cursorInvisible)
		if display("\\ (%0.2f s)\r", time.Since(st).Seconds(), c) {
			goto Wait
		}
		if display("| (%0.2f s)\r", time.Since(st).Seconds(), c) {
			goto Wait
		}
		if display("- (%0.2f s)\r", time.Since(st).Seconds(), c) {
			goto Wait
		}
		if display("/ (%0.2f s)\r", time.Since(st).Seconds(), c) {
			goto Wait
		}
	}
}
