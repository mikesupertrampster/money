package database

import (
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"os"
)

type DB struct {
	Client     *clover.DB
	dir        string
	Collection string
}

func Load() (DB, error) {
	db := DB{
		dir:        "_cache/db",
		Collection: "default",
	}

	if _, err := os.Stat(db.dir); os.IsNotExist(err) {
		if err = os.Mkdir(db.dir, 0755); err != nil {
			return db, err
		}
	}

	c, err := clover.Open(db.dir)
	if err != nil {
		return db, err
	}
	hasCollection, err := c.HasCollection(db.Collection)
	if err != nil {
		return db, err
	}
	if !hasCollection {
		if err = c.CreateCollection(db.Collection); err != nil {
			return db, err
		}
	}
	db.Client = c

	return db, nil
}

func (d *DB) Save(symbol string, i interface{}) error {
	doc := document.NewDocumentOf(i)

	all, err := d.Client.FindAll(d.Qry(symbol))
	if err != nil {
		return err
	}
	if len(all) == 0 {
		if err = d.Client.Insert(d.Collection, doc); err != nil {
			return err
		}
	} else {
		if err = d.Client.Update(d.Qry(symbol), doc.ToMap()); err != nil {
			return err
		}
	}
	if err = d.Client.Close(); err != nil {
		return err
	}

	return nil
}

func (d *DB) Qry(symbol string) *query.Query {
	return query.NewQuery(d.Collection).Where(query.Field("symbol").Eq(symbol))
}

func (d *DB) FindSymbol(symbol string) (*document.Document, error) {
	data, err := d.Client.FindFirst(d.Qry(symbol))
	if err != nil {
		return nil, err
	}
	return data, nil
}
