package main

import (
	"log"
	"money/types"
)

func run() error {
	//vantage, err := api.NewAlphaVantage()
	//if err != nil {
	//	return err
	//}

	//feeds, err := vantage.NewsSentiments([]string{"TSLA", "AAPL"})
	//if err != nil {
	//	return err
	//}
	//for _, feed := range feeds {
	//	err = feed.Save()
	//	if err != nil {
	//		return err
	//	}
	//}

	f := types.Feed{}
	all, err := f.GetAll()
	if err != nil {
		return err
	}
	for _, v := range all {
		println(v.Title)
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
