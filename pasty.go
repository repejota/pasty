package pasty

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"
)

// WatchClipboard runs forever watching the pastboard.
func WatchClipboard(ctx context.Context, out chan<- string, errs chan<- error) {
	for {
		text, err := ReadPasteBoard()
		if err != nil {
			errs <- err
		}

		if len(text) > 0 {
			out <- text
		}

		time.Sleep(2 * time.Second)
	}
}

// ReadPasteBoard reads the current value of the pastboard.
func ReadPasteBoard() (string, error) {
	cmd := exec.Command("pbpaste")
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// WritePasteBoard fills out the pasteboard with a string.
func WritePasteBoard(s string) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(s)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
