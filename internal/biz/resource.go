package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/sync/errgroup"
	"sync"
)

type ResourceDataInfo struct {
	ID        int64   `bson:"id"`
	CnName    string  `bson:"cnname"`
	EnName    string  `bson:"enname"`
	AliasName string  `bson:"aliasname"`
	Channel   string  `bson:"channel"`
	ChannelCN string  `bson:"channel_cn"`
	Area      string  `bson:"area"`
	ShowType  string  `bson:"show_type"`
	Expire    string  `bson:"expire"`
	Views     int64   `bson:"views"`
	Year      []int64 `bson:"year"`
}

type ResourceData struct {
	Info ResourceDataInfo `bson:"info"`
}

type Resource struct {
	ID        bson.ObjectID `bson:"_id"`
	CnName    string        `bson:"cnname"`
	EnName    string        `bson:"enname"`
	AliasName string        `bson:"aliasname"`
	Area      string        `bson:"area"`
	Views     int64         `bson:"views"`
	Data      ResourceData  `bson:"data"`
}

type ResourceRepo interface {
	ListTopByArea(context.Context, interface{}) (*[]Resource, error)
}

type ResourceUsecase struct {
	repo ResourceRepo
	log  *log.Helper
}

func NewResourceUsecase(repo ResourceRepo, logger log.Logger) *ResourceUsecase {
	return &ResourceUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ResourceUsecase) TopResourceMap(ctx context.Context) (map[string]*[]Resource, error) {
	areaFilter := map[string]bson.D{
		"ALL": {},
		"US":  {{"data.info.area", "美国"}},
		"JP":  {{"data.info.area", "日本"}},
		"KR":  {{"data.info.area", "韩国"}},
		"UK":  {{"data.info.area", "英国"}},
	}

	result := make(map[string]*[]Resource)
	var mu sync.Mutex // Mutex to protect the shared result map

	g, gCtx := errgroup.WithContext(ctx)

	for area, filter := range areaFilter {
		area, filter := area, filter // Capture loop variables
		g.Go(func() error {
			resources, err := uc.repo.ListTopByArea(gCtx, filter)
			if err != nil {
				return fmt.Errorf("failed to get top resources for %s: %w", area, err)
			}
			mu.Lock()
			result[area] = resources
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}
