package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/getlantern/systray"
	"github.com/repejota/pasty"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	systray.Run(pasty.OnReady, pasty.OnExit)
}
