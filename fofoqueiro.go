package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

import (
	"github.com/common-nighthawk/go-figure"
)

type ResourceStatus string

const (
	UP   ResourceStatus = "UP"
	DOWN ResourceStatus = "DOWN"
)

const version = 1.1
const monitoringTimes = 5
const monitoringDelay = 3
const logFileName = "monitoring.log"
const resourcesFileName = "resources.txt"

func main() {

	displayIntro()

	for {
		displayMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			readLogs()
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
	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println("Got error during recognizing command:", err)
	}

	return command
}

func startMonitoring() {
	urls := readResourcesToMonitoring()
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
	res, err := http.Get(resource)

	if err != nil {
		fmt.Println("Got error during checking resource", resource, ":", err)
	}

	if res.StatusCode == 200 {
		fmt.Println(resource, "seems healthy!")
		registerLog(resource, UP, res.StatusCode)
	} else {
		fmt.Println(resource, "seems unhealthy!. Got status code=", res.StatusCode)
		registerLog(resource, DOWN, res.StatusCode)
	}
}

func readResourcesToMonitoring() []string {
	var result []string

	file, err := os.Open(resourcesFileName)
	if err != nil {
		fmt.Println("Got error during reading resources file:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Got error during reading line from file:", err)
		}

		line = strings.TrimSpace(line)

		result = append(result, line)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Got error during close file:", err)
	}

	return result
}

func registerLog(resource string, status ResourceStatus, statusCode int) {
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Got error during interact with log file:", err)
	}

	_, err = file.WriteString("[" + time.Now().Format("2006-01-02 15:04:05") +
		"] " + resource + " - status(" + strconv.Itoa(statusCode) + "):" +
		string(status) + "\n")

	if err != nil {
		fmt.Println("Got error during writing file:", err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Got error during close file:", err)
	}
}

func readLogs() {
	file, err := ioutil.ReadFile(logFileName)

	if err != nil {
		fmt.Println("Got error during reading logs file:", err)
	}

	fmt.Println(string(file))
}
