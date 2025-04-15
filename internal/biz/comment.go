package biz

import (
	"context"
	"errors"

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
	Type       string        `bson:"type" json:"type"`
	ParentID   bson.ObjectID `bson:"parent_id" json:"parent_id"`
}

type CommentRepo interface {
	FindById(ctx context.Context, id string) (*Comment, error)
	ListLatest(ctx context.Context, size int64) ([]*Comment, error)
	List(ctx context.Context, resourceId, pageNo, pageSize int64, sort string) ([]*Comment, int64, error)
	Count(ctx context.Context) (int64, error)
	Search(ctx context.Context, keyword string) ([]*Comment, error)
	Save(ctx context.Context, comment *Comment) (*Comment, error)
}

type CommentUsecase struct {
	commentRepo  CommentRepo
	resourceRepo ResourceRepo
	log          *log.Helper
}

func NewCommentUsecase(commentRepo CommentRepo, resourceRepo ResourceRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{commentRepo: commentRepo, resourceRepo: resourceRepo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) ListLatest(ctx context.Context, size int64) ([]*Comment, error) {
	return uc.commentRepo.ListLatest(ctx, size)
}

func (uc *CommentUsecase) Count(ctx context.Context) (int64, error) {
	return uc.commentRepo.Count(ctx)
}

func (uc *CommentUsecase) Search(ctx context.Context, keyword string) ([]*Comment, error) {
	return uc.commentRepo.Search(ctx, keyword)
}

func (uc *CommentUsecase) List(ctx context.Context, resourceId, pageNo, pageSize int64, sort string) ([]*Comment, int64, error) {
	return uc.commentRepo.List(ctx, resourceId, pageNo, pageSize, sort)
}

func (uc *CommentUsecase) Create(ctx context.Context, comment *Comment) (*Comment, error) {
	resource, err := uc.resourceRepo.FindByID(ctx, comment.ResourceID)
	if err != nil {
		return nil, err
	}
	if resource == nil {
		return nil, errors.New("resource not found")
	}

	if !bson.ObjectID.IsZero(comment.ParentID) {
		parentComment, err := uc.commentRepo.FindById(ctx, comment.ParentID.Hex())
		if err != nil {
			return nil, err
		}
		if parentComment == nil {
			return nil, errors.New("parent comment not found")
		}
	}

	// todo: user-agent mw
	comment.Browser = "browser"
	return uc.commentRepo.Save(ctx, comment)
}
