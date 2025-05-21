package main

import (
	"log"
	"money/discord/bot"
	"os"
)

func run() error {
	//finviz, err := intelligence.NewFinviz()
	//if err != nil {
	//	return err
	//}
	//symbol, err := finviz.GetMetrics("CNEY")
	//if err != nil {
	//	return err
	//}
	//if err = symbol.Save(); err != nil {
	//	return err
	//}

	//var m types.Metrics
	//
	//db, err := database.Load()
	//if err != nil {
	//	return err
	//}
	//err = m.Get(db, "CNEY")
	//if err != nil {
	//	return err
	//}
	//println(m.Price)

	return nil
}

func main() {
	if err := bot.Start(os.Getenv("DISCORD_TOKEN")); err != nil {
		log.Fatal(err)
	}
	<-make(chan struct{})

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
