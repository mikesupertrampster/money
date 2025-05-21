package types

import (
	"github.com/ostafen/clover/v2/query"
	"money/database"
	"time"
)

type TitanPost struct {
	Rank      int         `clover:"rank"`
	Symbol    string      `clover:"symbol"`
	Heading   string      `clover:"heading"`
	Content   string      `clover:"content"`
	News      []TitanNews `clover:"news"`
	Image     string      `clover:"image"`
	TimeStamp time.Time   `clover:"timestamp"`
}

type TitanNews struct {
	Time     string `clover:"time"`
	Headline string `clover:"headline"`
	Link     string `clover:"link"`
}

func (t *TitanPost) Save() error {
	db, err := database.Load("titanposts")
	if err != nil {
		return err
	}
	if err = db.Save(t.Symbol, t); err != nil {
		return err
	}
	if err = db.Client.Close(); err != nil {
		return err
	}
	return nil
}

func (t *TitanPost) Get(symbol string) error {
	db, err := database.Load("titanposts")
	if err != nil {
		return err
	}
	data, err := db.FindSymbol(symbol)
	if err != nil {
		return err
	}
	if err = data.Unmarshal(&t); err != nil {
		return err
	}
	if err = db.Client.Close(); err != nil {
		return err
	}

	return nil
}

func (t *TitanPost) GetAll() ([]TitanPost, error) {
	all := make([]TitanPost, 0)

	db, err := database.Load("titanposts")
	if err != nil {
		return nil, err
	}
	posts, err := db.Client.FindAll(query.NewQuery("titanposts"))
	if err != nil {
		return nil, err
	}

	var post TitanPost
	for _, p := range posts {
		if err = p.Unmarshal(&post); err != nil {
			return nil, err
		}
		all = append(all, post)
	}

	if err = db.Client.Close(); err != nil {
		return nil, err
	}
	return all, err
}
