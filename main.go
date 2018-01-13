package main

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"os"
	"bufio"
	"regexp"
	"strings"
	"os/exec"
)

type thisTime struct {}

func main () {

	link := readInput()

	c := make(chan string)

	fmt.Println("Message: Yo! I'm watching, just sit back.")
	for _, link := range processInput(link) {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(30 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink (link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		var timeNow thisTime
		errorMsg := "["+ timeNow.getTime() +"]: " + link + " is down!\n"
		if _, err := os.Stat("log.txt"); err != nil {
			writeToLog(errorMsg)
		}
		file, _ := ioutil.ReadFile("log.txt")
		logData := string(file) + "\n " + errorMsg
		go func(link string) {
			notify(link)
		}(link)
		go writeToLog(logData)
		c <- link
		return
	}
	c <- link
	return
}

func (t thisTime) getTime() string {
	return time.Now().Format(time.UnixDate)
}

func writeToLog(data string) {
	ioutil.WriteFile(
		"log.txt",
		[]byte(data),
		0666,
	)
	return
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter website(s):  ")
	scanner.Scan()
	return scanner.Text()
}

func processInput(input string) []string {
	match, _ := regexp.MatchString(",", strings.Replace(input, " ", "", -1))
	if match {
		return strings.Split(input, ",")
	}
	return []string{input}
}

func notify(link string) {
	notification := fmt.Sprintf(
		"display notification \"%s\" with title \"%s\" subtitle \"%s\"",
		"Server is down",
		"Go App Up",
		"The link: " + link + " is down")
	exec.Command("osascript", "-e", notification).Run()
	return
}