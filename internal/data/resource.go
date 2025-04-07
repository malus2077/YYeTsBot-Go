package data

import (
	"YYeTsBot-Go/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type resourceRepo struct {
	data *Data
	log  *log.Helper
}

// NewResourceRepo .
func NewResourceRepo(data *Data, logger log.Logger) biz.ResourceRepo {
	return &resourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 新增文档结构体
type Info struct {
	Cnname string `bson:"cnname"`
}

type TData struct {
	Info Info `bson:"info"`
}

type ResourceModel struct {
	Data TData `bson:"data"`
}

func (r *resourceRepo) ListTopByArea(ctx context.Context, filter interface{}) (*[]biz.Resource, error) {
	opts := options.Find().
		SetProjection(bson.D{{"_id", 0}, {"data.info", 1}}).
		SetSort(bson.D{{"data.info.views", -1}}).
		SetLimit(15)

	cursor, err := r.data.db.Collection("yyets").Find(ctx, filter, opts)
	if err != nil {
		r.log.Errorf("Failed to find documents: %v", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var data []biz.Resource
	if err := cursor.All(ctx, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
