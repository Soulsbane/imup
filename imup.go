package main

import (
	"fmt"
	"time"

	"github.com/Soulsbane/imup/utils"
	"github.com/alexflint/go-arg"
	"github.com/imroc/req"
)

func keepFetchingPage(name string) {
	for {
		working := fetchPage(name)

		if working {
			second := utils.GetRandomSecond()
			time.Sleep(time.Duration(second) * time.Second)
		} /*else {
			time.Sleep(1 * time.Second)
		}*/
	}
}

func fetchPage(name string) bool {
	req.SetTimeout(3 * time.Second)
	_, err := req.Get(name)

	if err != nil {
		fmt.Print(utils.GetFormattedTime())
		fmt.Println(" - Reached timeout or error. Your internet is probably down!")

		return false
	}

	fmt.Print(utils.GetFormattedTime())
	fmt.Println(" - Your internet is working!")

	return true
}

func main() {
	var args struct {
		Repeat bool   `arg:"-r, --repeat" default:"false" help:"Keep checking periodically."`
		URL    string `arg:"-u, --url" default:"https://www.google.com/" help:"The website to fetch."`
	}

	arg.MustParse(&args)

	if args.Repeat {
		keepFetchingPage(args.URL)
	} else {
		fetchPage(args.URL)
	}
}
