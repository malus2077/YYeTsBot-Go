package data

import (
	"YYeTsBot-Go/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (r *commentRepo) FindById(ctx context.Context, id string) (*biz.Comment, error) {
	collection := r.data.db.Collection("comment")
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	condition := bson.M{
		"_id": objectId,
	}
	var comment *biz.Comment
	err = collection.FindOne(ctx, condition).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepo) Save(ctx context.Context, comment *biz.Comment) (*biz.Comment, error) {
	collection := r.data.db.Collection("comment")
	result, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment.ID = result.InsertedID.(bson.ObjectID)
	return comment, nil
}

func (r *commentRepo) List(ctx context.Context, resourceId, pageNo, pageSize int64, sort string) ([]*biz.Comment, int64, error) {
	sorting := -1
	if sort == "oldest" {
		sorting = 1
	}

	collection := r.data.db.Collection("comment")
	condition := bson.M{
		"resource_id": resourceId,
		"deleted_at":  bson.M{"$exists": false},
		"type":        bson.M{"$ne": "child"},
	}
	projection := bson.M{
		"ip":        0,
		"parent_id": 0,
	}
	opts := options.Find().
		SetProjection(projection).
		SetSort(bson.M{"_id": sorting}).
		SetSkip((pageNo - 1) * pageSize).
		SetLimit(pageSize)

	count, err := collection.CountDocuments(ctx, condition)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, condition, opts)
	if err != nil {
		return nil, count, err
	}
	defer cursor.Close(ctx)

	var comments []*biz.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, count, err
	}
	return comments, count, nil
}

func (r *commentRepo) Search(ctx context.Context, keyword string) ([]*biz.Comment, error) {
	collection := r.data.db.Collection("comment")
	condition := bson.M{
		"content": bson.M{
			"$regex":   keyword,
			"$options": "i",
		},
	}

	// 复用通用查询参数
	projection := bson.M{
		"ip":        0,
		"parent_id": 0,
		"children":  0,
	}

	// 构建查询选项
	opts := options.Find().
		SetProjection(projection).
		SetSort(bson.D{{"_id", -1}})

	cursor, err := collection.Find(ctx, condition, opts)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	var data []*biz.Comment
	if err = cursor.All(ctx, &data); err != nil {
		r.log.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	return data, nil
}

func (r *commentRepo) ListLatest(ctx context.Context, size int64) ([]*biz.Comment, error) {
	condition := bson.M{"deleted_at": bson.M{"$exists": false}}

	projection := bson.M{
		"ip":        0,
		"parent_id": 0,
		"children":  0,
	}

	opts := options.Find().
		SetProjection(projection).
		SetSort(bson.D{{Key: "_id", Value: -1}}).
		SetLimit(size)

	cursor, err := r.data.db.Collection("comment").Find(ctx, condition, opts)
	if err != nil {
		return nil, err
	}

	var data []*biz.Comment
	if err = cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return data, nil
}

func (r *commentRepo) Count(ctx context.Context) (int64, error) {
	condition := bson.M{"deleted_at": bson.M{"$exists": false}}
	count, err := r.data.db.Collection("comment").CountDocuments(ctx, condition)
	return count, err
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
