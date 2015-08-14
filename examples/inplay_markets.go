package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/gpavlidi/gobet"
)

func main() {
	username := flag.String("u", os.Getenv("BETFAIR_USERNAME"), "Betfair username")
	password := flag.String("p", os.Getenv("BETFAIR_PASSWORD"), "Betfair password")
	appKey := flag.String("key", os.Getenv("BETFAIR_APPKEY"), "Betfair appkey")

	flag.Parse()

	if *username == "" || *password == "" {
		flag.Usage()
	}

	conf := gobet.Configuration{Username: *username, Password: *password, AppKey: *appKey, Locale: "en", LogDest: os.Stdout}
	s, err := gobet.NewSession(&conf)
	if err != nil {
		log.Fatal("Invalid Configuration! Exiting...")
	}

	err = s.Login()
	if err != nil {
		log.Fatal("Could not login! Exiting...")
	}
	defer s.Logout()

	log.Println("Setting Periodic Keep Alive")
	s.DoKeepAliveEvery(time.Hour * 4)

	filter := gobet.MarketFilter{TurnInPlayEnabled: true, MarketBettingTypes: []gobet.MarketBettingType{"ODDS"}}
	input := gobet.BettingInputParams{Filter: &filter}
	events, err := s.ListEvents(input)

	for _, event := range events {
		log.Println(event.Event.Name)
	}
}
