package main

import (
	"fmt"

	"github.com/saadtresayyed10/learn-go-discord/bot"
	"github.com/saadtresayyed10/learn-go-discord/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
