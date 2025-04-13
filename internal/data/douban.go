package data

import (
	"YYeTsBot-Go/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type doubanRepo struct {
	data *Data
	log  *log.Helper
}

func NewDoubanRepo(data *Data, logger log.Logger) biz.DoubanRepo {
	return &doubanRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *doubanRepo) FindByResourceId(ctx context.Context, resourceId int64) (*biz.Douban, error) {
	collection := r.data.db.Collection("douban")
	filter := bson.M{"resourceId": resourceId}
	opts := options.FindOne().SetProjection(bson.M{
		"_id": 0,
		"raw": 0,
	})
	var douban *biz.Douban
	err := collection.FindOne(ctx, filter, opts).Decode(&douban)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return douban, nil
}

func (r *doubanRepo) Insert(ctx context.Context, resourceName string, resourceId int64) (*biz.Douban, error) {
	//todo: 性能优化
	douban := r.data.crawler.Douban.FetchMediaInfo(resourceName)
	douban.ResourceId = resourceId

	if douban == nil {
		return nil, errors.New("douban info not found")
	}

	collection := r.data.db.Collection("douban")
	//todo: context timeout
	result, err := collection.InsertOne(context.Background(), douban)
	if err != nil {
		return nil, err
	}
	douban.Id = result.InsertedID.(bson.ObjectID)
	return douban, nil
}
