package main

import (
	"encoding/json"
	// "fmt"
	"time"

	N "github.com/NullpointerW/anicat/net"
	"github.com/gosuri/uiprogress"
)

type TorrentProgressListSend struct {
	List []TorrentProgress `json:"list"`
	Fin  bool              `json:"fin"`
}
type TorrentProgress struct {
	Percentage int    `json:"percentage"`
	Name       string `json:"name"`
}

func statsBuiltinProcess(conn *N.Conn) error {
	bars := make(map[string]*uiprogress.Bar)
	uiprogress.Start()
	for {
		b, err := conn.Read()
		if err != nil {
			return err
		}
		l := new(TorrentProgressListSend)
		err = json.Unmarshal(b, l)
		if err != nil {
			return err
		}
		fin := l.Fin
		for _, t := range l.List {
			if bar, ex := bars[t.Name]; !ex {
				nbar := uiprogress.AddBar(100)
				tname:=t.Name
				nbar.PrependFunc(func(b *uiprogress.Bar) string {
					return tname
				})
				nbar.AppendCompleted()
				_ = nbar.Set(t.Percentage)
				bars[t.Name] = nbar
			} else {
				bar.Set(t.Percentage)
			}
		}
		if fin {
			time.Sleep(1 * time.Second)
			return nil
		}
	}
}
