package main

import (
	"log"
	"money/intelligence/scrape"
)

func run() error {
	s, err := intelligence.NewStockTwits()
	if err != nil {
		return err
	}
	err = s.GetSentiments("AAPL")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//if err := bot.Start(os.Getenv("DISCORD_TOKEN")); err != nil {
	//	log.Fatal(err)
	//}
	//<-make(chan struct{})
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
