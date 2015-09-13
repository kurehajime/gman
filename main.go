// main.go
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	open bool
	list bool
)

func main() {
	flag.BoolVar(&open, "o", false, "Open web browser")
	flag.BoolVar(&list, "l", false, "List  repository")
	flag.Parse()
	if len(flag.Args()) == 0 {
		return
	}
	text := flag.Args()[0]

	var res string
	var err error
	switch true {
	case open == true:
		res, err = OpenRepo(text)
	case list == true:
		res, err = ShowList(text)
	default:
		res, err = Gman(text)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Print(res)
}
