package main

import (
	"log"
)

func run() error {
	//s, err := intelligence.NewStockTitans()
	//if err != nil {
	//	return err
	//}
	//news, err := s.GetNews()
	//if err != nil {
	//	return err
	//}
	//
	//for _, n := range news {
	//	if err = n.Save(); err != nil {
	//		return err
	//	}
	//}

	//var m types.TitanPost
	//all, err := m.GetAll()
	//if err != nil {
	//	return err
	//}
	//println(len(all))

	//if err = symbol.Save(); err != nil {
	//	return err
	//}

	//var m types.Metrics
	//if err := m.Get("CNEY"); err != nil {
	//	return err
	//}
	//println(m.Price)

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
