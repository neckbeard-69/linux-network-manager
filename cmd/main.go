package main

import (
	"log"
	"nmcli-tui/pkg/ui"
	"nmcli-tui/pkg/utils"
)

func main() {
	utils.ClearScr()
	choice, err := ui.GetUserChoice()
	if err != nil {
		log.Fatal(err)
	}
	ui.Action(choice)
}
