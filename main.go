package main

import (
	doc "./docread"
	fetched "./fetchword"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("\nargs not right.\nUsage:\n\t ./fetch path")
	}
	path := args[0]
	s := doc.Read(path)
	sl := strings.Fields(s)
	wl := fetched.GetWords(sl)
	for _, wd := range wl {
		fmt.Println(wd)
	}
}

