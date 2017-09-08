package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/getlantern/systray"
)

var (
	// Version is the latest release number.
	//
	// This number is the latest tag from the git repository.
	Version string

	// Build is the lastest build string.
	//
	// This string is the branch name and the commit hash (short format)
	Build string
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("pasty: ")

	var (
		// Define CLI flags.
		versionFlag = flag.Bool("version", false, "Show version information")
	)

	flag.Parse()
	if *versionFlag {
		versionInfo := showVersion()
		fmt.Println(versionInfo)
		os.Exit(0)
	}

	fmt.Println("onStart")
	systray.Run(onReady, onExit)
}

func onExit() {
	fmt.Println("onExit")
}

func onReady() {
	fmt.Println("onReady")
	systray.SetTitle("P")
	systray.SetTooltip("Pasty")
	go monitorClipboard()
}

// monitorClipboard mo
func monitorClipboard() {
	for {
		text, err := readPasteBoard()
		if err != nil {
			log.Println("error", err)
		}
		log.Println(text)
		time.Sleep(2 * time.Second)
	}
}

// readClipboard reads the current value of the pastboard.
func readPasteBoard() (string, error) {
	command := exec.Command("pbpaste")
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// showVersion prints the current version information.
func showVersion() string {
	versionInfo := fmt.Sprintf("pasty : Version %s Build %s", Version, Build)
	return versionInfo
}
