package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

import (
	"github.com/common-nighthawk/go-figure"
)

const monitoringTimes = 5
const monitoringDelay = 3

func main() {

	displayIntro()

	for {
		displayMenu()
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
	const version = 1.1
	myFigure := figure.NewFigure("Fofoqueiro", "", true)
	myFigure.Print()

	fmt.Println("Version", version)
}

func displayMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit")
}

func readCommand() int {
	var command int
	_, _ = fmt.Scan(&command)

	return command
}

func startMonitoring() {
	urls := []string{"https://random-status-code.herokuapp.com/", "https://www.diegoramos.me/"}
	fmt.Println("Monitoring", len(urls), "resources")

	for index := 0; index < monitoringTimes; index++ {
		for _, resource := range urls {
			checkResource(resource)
		}
		time.Sleep(monitoringDelay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func checkResource(resource string) {
	res, _ := http.Get(resource)
	if res.StatusCode == 200 {
		fmt.Println(resource, "seems healthy!")
	} else {
		fmt.Println(resource, "seems unhealthy!. Got status code=", res.StatusCode)
	}
}
