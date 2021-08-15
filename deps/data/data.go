package data

import (
	"log"
	"smallhouse/ent"

	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var ProvideSet = wire.NewSet(NewEntClient, NewData)

type Data struct {
	db *ent.Client
}

func NewEntClient() *ent.Client {
	client, err := ent.Open(
		"mysql",
		"user:q123q123@master/nomad",
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	return client
}

func NewData(entClient *ent.Client) (*Data, func(), error) {
	d := &Data{
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Println(err)
		}
	}, nil
}
