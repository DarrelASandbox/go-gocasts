package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel) // use go keyword only in front of function call for goroutine
	}

	// Blocking channels
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }

	// for {
	// 	go checkLink(channel, channel)
	// }

	// for link := range channel {
	// 	go checkLink(link, channel)
	// }

	// for link := range channel {
	// time.Sleep(5 * time.Second)
	// 	go checkLink(link, channel)
	// }

	// function literal channel gotcha
	// for link := range channel {
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(link, channel)
	// 	}()
	// }

	// Fix for gotcha
	for link := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "cannot be reached.")
		c <- "Might be down."
		return
	}

	fmt.Println(link, "is OK!")
	c <- link
}
