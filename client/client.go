/*
HTTP Client Example
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func doRequests() {
	ch := make(chan bool, 3)
	for i := 0; i < 10; i++ {
		go func(id int) {
			ch <- true
			strID := strconv.Itoa(id)
			res, err := http.PostForm("http://localhost:8080", url.Values{"key": {"Value"}, "id": {strID}})
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			<-ch
		}(i)
	}
	for i := 0; i < 3; i++ {
		ch <- true
	}
}

func noobRequests() {
	for i := 0; i < 10; i++ {
		func(id int) {
			strID := strconv.Itoa(id)
			res, err := http.PostForm("http://localhost:8080", url.Values{"key": {"Value"}, "id": {strID}})
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
		}(i)
	}
}

func main() {
	reading := true
	for reading {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		switch text {
		case "call":
			fmt.Println("Making Requests")
			doRequests()
		case "quit":
			reading = false
		case "noob":
			fmt.Println("Making noob requests")
			noobRequests()
		default:
			fmt.Println("Command not supported")
		}
	}
}
