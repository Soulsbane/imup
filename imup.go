package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/imroc/req"
)

func getRandomSecond() int {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 60

	return rand.Intn(max-min+1) + min
}

func fetchPage(name string) bool {

	req.SetTimeout(5 * time.Second)
	_, err := req.Get(name)

	if err != nil {
		fmt.Println("Reached timeout or error. Your internet is probably down!")
		return false
	}

	fmt.Println("Your internet is working!")
	return true
}

func main() {
	fetchPage("http://google.com/")
}
