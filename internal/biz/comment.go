package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Comment struct {
	ID         bson.ObjectID `bson:"_id" json:"id,omitempty"`
	Date       string        `bson:"date" json:"date"`
	Content    string        `bson:"content" json:"content"`
	ResourceID int64         `bson:"resource_id" json:"resource_id"`
	Browser    string        `bson:"browser" json:"browser"`
	Username   string        `bson:"username" json:"username"`
}

type CommentRepo interface {
	ListLatest(ctx context.Context, size int64) ([]*Comment, error)
	List(ctx context.Context, resourceId, pageNo, pageSize int64, sort string) ([]*Comment, int64, error)
	Count(ctx context.Context) (int64, error)
	Search(ctx context.Context, keyword string) ([]*Comment, error)
}

type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) ListLatest(ctx context.Context, size int64) ([]*Comment, error) {
	return uc.repo.ListLatest(ctx, size)
}

func (uc *CommentUsecase) Count(ctx context.Context) (int64, error) {
	return uc.repo.Count(ctx)
}

func (uc *CommentUsecase) Search(ctx context.Context, keyword string) ([]*Comment, error) {
	return uc.repo.Search(ctx, keyword)
}

func (uc *CommentUsecase) List(ctx context.Context, resourceId, pageNo, pageSize int64, sort string) ([]*Comment, int64, error) {
	return uc.repo.List(ctx, resourceId, pageNo, pageSize, sort)
}
