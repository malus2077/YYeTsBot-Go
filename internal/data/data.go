package data

import (
	"YYeTsBot-Go/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/eko/gocache/lib/v4/cache"
	redisStore "github.com/eko/gocache/store/redis/v4"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db          *mongo.Database
	cacheManger *cache.Cache[string]
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	client, err := mongo.Connect(options.Client().ApplyURI(c.Database.Source))
	if err != nil {
		return nil, nil, err
	}

	db := client.Database(c.Database.DbName)

	store := redisStore.NewRedis(redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	}))

	cacheManger := cache.New[string](store)

	cleanup := func() {
		log.Info("closing the data resources")
		if err := client.Disconnect(context.Background()); err != nil {
			log.Error(err)
		}
	}
	return &Data{db: db, cacheManger: cacheManger}, cleanup, nil
}
