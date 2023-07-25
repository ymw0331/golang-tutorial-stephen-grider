package main

import (
	"fmt"
	"net/http"
	"time"
)

// status checker --> http request  --> http://google.com / http://facebook.com / http://stackoverflow.com
func main() {

	// list out url of each websites inside of a slice of type string
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stakeoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string) //make is a built in function to create value out of a given type

	// loop through slice, every url inside slice make a http request
	for _, link := range links {
		go checkLink(link, c)
	}

	//1st way
	// for i := 0; i < len(links); i++ {  //convert to one below
	// 	fmt.Println(<-c)
	// }

	//2nd way
	// for {
	// 	go checkLink(<-c, c)
	// }

	//3rd way
	for l := range c {
		go checkLink(l, c)
	}

	for l := range c {
		go checkLink(l, c)
		go func(link string) { //function literal
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l) //l as function argument
	}

	// fmt.Println(<-c) // main routine
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c) // stuck here because no 6th link

}

func checkLink(link string, c chan string) {

	time.Sleep(5 * time.Second)

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up" // go routine
}

// Equivalent of Funtion Literal
// Go (FL) <-> JS (Anonymouse Function) <-> Ruby (Lambda) <-> Python(Lamdha)
