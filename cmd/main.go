package main

import (
	"github.com/neckbeard-69/linux-network-manager/pkg/ui"
	"github.com/neckbeard-69/linux-network-manager/pkg/utils"
	"log"
)

func main() {
	utils.ClearScr()
	choice, err := ui.GetUserChoice()
	if err != nil {
		log.Fatal(err)
	}
	ui.Action(choice)
}
