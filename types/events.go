package types

import (
	"github.com/ostafen/clover/v2/query"
	"money/database"
	"time"
)

type AiEvent struct {
	Symbol    string    `clover:"symbol"`
	Title     string    `clover:"title"`
	Sentiment string    `clover:"sentiment"`
	Content   string    `clover:"content"`
	Link      string    `clover:"link"`
	Image     string    `clover:"image"`
	TimeStamp time.Time `clover:"timestamp"`
}

func (e *AiEvent) Save() error {
	db, err := database.Load("aievents")
	if err != nil {
		return err
	}
	if err = db.Save(e.Symbol, e); err != nil {
		return err
	}
	if err = db.Client.Close(); err != nil {
		return err
	}
	return nil
}

func (e *AiEvent) GetAll() ([]AiEvent, error) {
	all := make([]AiEvent, 0)

	db, err := database.Load("aievents")
	if err != nil {
		return nil, err
	}
	posts, err := db.Client.FindAll(query.NewQuery("aievents"))
	if err != nil {
		return nil, err
	}

	var event AiEvent
	for _, p := range posts {
		if err = p.Unmarshal(&event); err != nil {
			return nil, err
		}
		all = append(all, event)
	}

	if err = db.Client.Close(); err != nil {
		return nil, err
	}
	return all, err
}
