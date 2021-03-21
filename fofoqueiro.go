package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	displayIntro()
	displayMenu()

	for {
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Quitting...")
			os.Exit(0)
		default:
			fmt.Println("Unrecognized command!")
			os.Exit(-1)
		}
	}

}

func displayIntro() {
	const name = "Ramos"
	const version = 1.1
	fmt.Println("Hello sir,", name)
	fmt.Println("Version", version)
}

func displayMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("You chose", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://random-status-code.herokuapp.com/"

	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		fmt.Println(site, "seems healthy!")
	} else {
		fmt.Println(site, "seems unhealthy!. Got status code=", res.StatusCode)
	}
}
