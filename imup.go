package main

import (
	"fmt"
	"time"

	"github.com/imroc/req"
)

func fetchPage(name string) bool {

	req.SetTimeout(5 * time.Second)
	_, err := req.Get(name)

	if err != nil {
		//log.Fatal(err)
		fmt.Println("Reached timeout or error. Your internet is probably down!")
		return false
	}

	fmt.Println("Your internet is working!")
	return true
}

func main() {
	fetchPage("http://google.com/")
}
