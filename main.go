package main

import "fmt"

func main() {
	fmt.Println("onStart")

	systray.Run(onReady, onExit)
}

func onExit() {
	fmt.Println("onExit")
}

func onReady() {
	fmt.Println("onReady")
}
