package pasty

import (
	"context"
	"log"
	"strings"

	"github.com/getlantern/systray"
)

// OnExit ...
func OnExit() {
	log.Println("onExit")
}

// OnReady ...
func OnReady() {
	systray.SetTitle("Pb")
	systray.SetTooltip("Pasteboard Tooltip")

	// main logic

	var item string
	var err error
	var size, idx int

	size = 20
	idx = 0
	store := make([]string, size)
	menu := make([]*systray.MenuItem, size)
	texts := make(chan string)
	errs := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go WatchClipboard(ctx, texts, errs)

	for {
		select {
		case <-ctx.Done():
			if err = ctx.Err(); err != nil {
				log.Fatal(err)
				return
			}
			log.Println("ctx done")
			return
		case err = <-errs:
			if err != nil {
				return
			}
		case item = <-texts:
			if uniqueText(item, store) {
				store[idx] = item
				addItemToMenu(idx, store, menu)
				handleIndex(&idx, size)
			}
		}
	}
}

func addItemToMenu(i int, store []string, menu []*systray.MenuItem) {
	if menu[i] == nil {
		menu[i] = createTrayBtn(store[i])
		go listenMenuChecked(i, menu[i], store[i])
	} else {
		menu[i].SetTitle(getTitle(store[i]))
		menu[i].SetTooltip(store[i])
		go listenMenuChecked(i, menu[i], store[i])
	}
}

func listenMenuChecked(i int, menuItem *systray.MenuItem, text string) {
	select {
	case <-menuItem.ClickedCh:
		WritePasteBoard(text)
		go listenMenuChecked(i, menuItem, text)
		return
	}
}

func getTitle(item string) (title string) {
	if len(item) > 20 {
		title = item[:20] + "..."
	} else {
		title = item
	}
	title = strings.TrimSpace(title)
	title = strings.Replace(title, "\n", " ", -1)
	return
}

func createTrayBtn(item string) *systray.MenuItem {
	return systray.AddMenuItem(getTitle(item), item)
}

// uniqueText checks if the value of s is in the store.
func uniqueText(s string, store []string) bool {
	for i := range store {
		if store[i] == s {
			return false
		}
	}
	return true
}
