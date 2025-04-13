package data

import (
	"YYeTsBot-Go/internal/conf"
	"YYeTsBot-Go/internal/data/crawler"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/eko/gocache/lib/v4/cache"
	redisStore "github.com/eko/gocache/store/redis/v4"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewGreeterRepo, NewResourceRepo, NewCommentRepo, NewDoubanRepo, NewCaptchaRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db           *mongo.Database
	cacheManager *cache.Cache[string]
	crawler      *crawler.Crawler
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	hLog := log.NewHelper(logger)
	client, err := mongo.Connect(options.Client().ApplyURI(c.Database.Source).SetConnectTimeout(5 * time.Second))
	if err != nil {
		return nil, nil, err
	}

	db := client.Database(c.Database.DbName)

	if err := db.Client().Ping(context.Background(), nil); err == nil {
		hLog.Infof("Successfully connected to MongoDB!")
	}

	store := redisStore.NewRedis(redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	}))

	cacheManager := cache.New[string](store)
	crawler := &crawler.Crawler{}

	cleanup := func() {
		hLog.Info("closing the data resources")
		if err := client.Disconnect(context.Background()); err != nil {
			hLog.Error(err)
		}
	}
	return &Data{db: db, cacheManager: cacheManager, crawler: crawler}, cleanup, nil
}
