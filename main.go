package main

import (
	"log"
	"github.com/yanzay/tbot"
	"strconv"
	"time"
	"github.com/ciazhar/config"
)


func main() {
	conf := config.Load()

	bot, err := tbot.NewServer(conf.Get("telegramToken").String())
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/answer", "42")
	bot.HandleFunc("/timer {seconds}", timerHandler)
	bot.ListenAndServe()
}

func timerHandler(m *tbot.Message) {
	// m.Vars contains all variables, parsed during routing
	secondsStr := m.Vars["seconds"]
	// Convert string variable to integer seconds value
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		m.Reply("Invalid number of seconds")
		return
	}
	m.Replyf("Timer for %d seconds started", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}
