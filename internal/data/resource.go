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

type resourceRepo struct {
	data *Data
	log  *log.Helper
}

func (r *resourceRepo) FindByID(ctx context.Context, id int64) (*biz.Resource, error) {
	collection := r.data.db.Collection("yyets")
	filter := bson.M{"data.info.id": id}
	update := bson.M{"$inc": bson.M{"data.info.views": 1}}
	opts := options.FindOneAndUpdate().
		SetProjection(bson.M{"_id": false}).
		SetReturnDocument(options.After)

	var resource biz.Resource
	if err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&resource); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &resource, nil
}

func (r *resourceRepo) FetchMapByIds(ctx context.Context, ids []int64) (map[int64]*biz.Resource, error) {
	collection := r.data.db.Collection("yyets")
	filter := bson.M{"data.info.id": bson.M{"$in": ids}}
	opts := options.Find().SetProjection(bson.M{"data.info": 1})
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var res []*biz.Resource
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*biz.Resource)
	for _, item := range res {
		result[item.Data.Info.ID] = item
	}
	return result, nil
}

func (r *resourceRepo) Search(ctx context.Context, keyword string) ([]*biz.Resource, error) {
	collection := r.data.db.Collection("yyets")

	filter := bson.M{
		"$or": []bson.M{
			{"data.info.cnname": bson.M{"$regex": keyword, "$options": "i"}},
			{"data.info.enname": bson.M{"$regex": keyword, "$options": "i"}},
			{"data.info.aliasname": bson.M{"$regex": keyword, "$options": "i"}},
		},
	}

	opts := options.Find().
		SetProjection(bson.M{
			"data.info": 1,
		}).
		SetLimit(100)

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var resources []*biz.Resource
	if err := cursor.All(ctx, &resources); err != nil {
		return nil, err
	}

	return resources, nil
}

func (r *resourceRepo) ListTopByArea(ctx context.Context, filter interface{}) (*[]biz.Resource, error) {
	opts := options.Find().
		SetProjection(bson.D{{"_id", 0}, {"data.info", 1}}).
		SetSort(bson.D{{"data.info.views", -1}}).
		SetLimit(15)

	cursor, err := r.data.db.Collection("yyets").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var data []biz.Resource
	if err := cursor.All(ctx, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func NewResourceRepo(data *Data, logger log.Logger) biz.ResourceRepo {
	return &resourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
