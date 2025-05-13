package ui

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/neckbeard-69/linux-network-manager/pkg/utils"
)

func GetUserChoice() (int, error) {
	fmt.Println("Please choose what you wanna do:")
	fmt.Println("1. connect to a wifi network")
	fmt.Println("2. disconnect from the current network")
	fmt.Print("----------------------\n")
	fmt.Print("Write the number curresponding to your choice: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("there was an error getting the input\n")
		return -1, err
	}
	choice, err := strconv.Atoi(text[0:1])
	if err != nil {
		log.Fatal("Your input must be an integer!")
		return -1, err
	}
	if choice < 1 || choice > 2 {
		log.Fatal("Please provide a valid option")
		return -1, errors.New("Invalid option")
	}

	return choice, nil
}

func Action(choice int) error {
	switch choice {
	case 1:
		utils.ClearScr()
		nets, err := utils.GetAvailableNetworks()
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("Please pick a network to connect to:")
		fmt.Print("\n")
		for i, net := range nets {
			fmt.Printf("%d. %s\n", i+1, net)
		}
		fmt.Print("----------------------\n")
		fmt.Print("Write the number of the desired Wi-Fi network: ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')

		chosenNet, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			log.Fatal(err)
		}
		if chosenNet-1 < 0 || chosenNet > len(nets) {
			log.Fatal(errors.New("Invalid network"))
		}
		err = utils.Connect(nets[chosenNet-1])
		if err != nil {
			log.Fatal(err)
		}
		break
	case 2:
		err := utils.DisconnectFromCurrentNetwork()
		if err != nil {
			log.Fatal(err)
		}
	default:
		return errors.New("Invalid choice")
	}
	return nil
}
