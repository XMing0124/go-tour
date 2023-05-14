package main

import (
	"log"
	"github.com/go-programming-tour-book/tour/cmd"
)

func main() {
	err := cmd.Exeute()
	if err!= nil {
		log.Fatal("command execute error: ",err)
	}
}
