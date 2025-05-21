package api

import (
	"context"
	"github.com/polygon-io/client-go/rest/iter"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

type Polygon struct {
	rateLimit string
	client    *polygon.Client
}

func NewPolygon() (*Polygon, error) {
	return &Polygon{
		rateLimit: "5 API Calls / Minute",
		client:    polygon.New(os.Getenv("POLYGON_KEY")),
	}, nil
}

func (t *Polygon) ShortInterest(symbol string) (*iter.Iter[models.ShortInterest], error) {
	params := models.ListShortInterestParams{}.
		WithTicker(models.EQ, symbol).
		WithOrder(models.Order("asc")).
		WithLimit(10).
		WithSort(models.Sort("ticker"))

	itr := t.client.VX.ListShortInterest(context.Background(), params)
	if itr.Err() != nil {
		return nil, itr.Err()
	}

	return itr, nil
}

func (t *Polygon) ShortVolume(symbol string) (*iter.Iter[models.ShortVolume], error) {
	params := models.ListShortVolumeParams{}.
		WithTicker(models.EQ, symbol).
		WithOrder(models.Order("asc")).
		WithLimit(10).
		WithSort(models.Sort("ticker"))

	itr := t.client.VX.ListShortVolume(context.Background(), params)
	if itr.Err() != nil {
		return nil, itr.Err()
	}

	return itr, nil
}
