package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/imroc/req"
)

func getRandomSecond() int {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 60

	return rand.Intn(max-min+1) + min
}

func keepFetchingPage(name string) {
	for {
		second := getRandomSecond()
		fetchPage(name)
		time.Sleep(time.Duration(second) * time.Second)
	}
}

func fetchPage(name string) bool {

	req.SetTimeout(5 * time.Second)
	_, err := req.Get(name)

	// TODO: Add time logging.
	if err != nil {
		fmt.Println("Reached timeout or error. Your internet is probably down!")
		return false
	}

	fmt.Println("Your internet is working!")
	return true
}

func main() {
	var args struct {
		Repeat bool `arg:"-r, --repeat" default:"false" help:"Keep checking periodically."`
	}

	arg.MustParse(&args)

	if args.Repeat {
		keepFetchingPage("http://google.com/")
	} else {
		fetchPage("http://google.com/")
	}
}
