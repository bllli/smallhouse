package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"smallhouse/ent"
	"smallhouse/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	client *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to db: %v", err)
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		err := client.Close()
		if err != nil {
			log.NewHelper(logger).Errorf("clos data resource fail %v", err)
		}
	}
	return &Data{
		client,
	}, cleanup, nil
}
