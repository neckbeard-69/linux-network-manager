package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ConnectWithPassword(SSID, password string) error {
	cmd := exec.Command("nmcli", "device", "wifi", "connect", SSID, "password", password)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(string(output))
	return nil
}

func deleteNetwork(SSID string) error {

	cmd := exec.Command("nmcli", "delete", "delete", SSID)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(string(output))
	return nil
}
func ConnectWithoutPassword(SSID string) error {
	cmd := exec.Command("nmcli", "connection", "up", SSID)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(string(output))
	return nil
}

func DisconnectFromCurrentNetwork() error {
	cmd := exec.Command("nmcli", "-t", "-f", "NAME", "connection", "show", "--active")
	network, err := cmd.Output()
	if err != nil {
		log.Println("Error executing nmcli to fetch active network:", err)
		return fmt.Errorf("failed to fetch active network: %v", err)
	}

	networkName := strings.TrimSpace(string(network))
	lines := strings.Split(string(networkName), "\n")
	for _, line := range lines {
		if !strings.Contains(strings.ToLower(line), "wired") {
			networkName = line
			break
		}
	}
	if networkName == "" {
		return fmt.Errorf("no active network found")
	}
	log.Println("Active network:", networkName)

	cmd = exec.Command("nmcli", "connection", "down", networkName)
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error executing nmcli to disconnect from network:", err)
		return fmt.Errorf("failed to disconnect from network '%s': %v", networkName, err)
	}
	fmt.Println(string(output))
	return nil
}

func GetAvailableNetworks() ([]string, error) {
	cmd := exec.Command("nmcli", "-f", "SSID", "device", "wifi", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	// Skip the header line
	lines = lines[1:]

	var networks []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			networks = append(networks, line)
		}
	}

	return networks, nil
}

func IsNetworkSaved(SSID string) bool {
	cmd := exec.Command("nmcli", "-t", "-f", "NAME", "connection", "show")
	output, _ := cmd.Output()

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == SSID {
			return true
		}
	}
	return false
}

func IsNetworkOpen(SSID string) bool {
	cmd := exec.Command("nmcli", "-t", "-f", "SSID,SECURITY", "device", "wifi", "list")
	output, _ := cmd.Output()

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) >= 2 && fields[0] == SSID {
			return fields[1] == ""
		}
	}
	return false
}

func Connect(SSID string) error {
	if IsNetworkSaved(SSID) {
		err := ConnectWithoutPassword(SSID)
		if err == nil {
			return nil
		}
		log.Println("connection failed")
	}

	fmt.Print("Enter the password for the network: ")
	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	deleteNetwork(SSID)
	return ConnectWithPassword(SSID, strings.TrimSpace(password))
}

func ClearScr() {
	fmt.Print("\033[H\033[2J")
}
