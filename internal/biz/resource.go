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

type ResourceFile struct {
	Way     string `bson:"way" json:"way"`
	WayCn   string `bson:"way_cn" json:"way_cn"`
	Address string `bson:"address" json:"address"`
	Passwd  string `bson:"passwd" json:"passwd"`
}

type ResourceSeasonItem struct {
	ItemId     string         `bson:"itemid" json:"itemid"`
	Episode    string         `bson:"episode" json:"episode"`
	Name       string         `bson:"name" json:"name"`
	Size       string         `bson:"size" json:"size"`
	YyetsTrans int64          `bson:"yyets_trans" json:"yyets_trans"`
	Dateline   string         `bson:"dateline" json:"dateline"`
	Files      []ResourceFile `bson:"files" json:"files"`
}

type ResourceSeason struct {
	SeasonNum string                           `bson:"season_num" json:"season_num"`
	SeasonCn  string                           `bson:"season_cn" json:"season_cn"`
	Items     map[string][]*ResourceSeasonItem `bson:"items" json:"items"`
	Formats   []string                         `bson:"formats" json:"formats"`
}

type ResourceData struct {
	Info ResourceDataInfo  `bson:"info"`
	List []*ResourceSeason `bson:"list"`
}

type Resource struct {
	Id        int64        `bson:"id"`
	CnName    string       `bson:"cnname"`
	EnName    string       `bson:"enname"`
	AliasName string       `bson:"aliasname"`
	Area      string       `bson:"area"`
	Views     int64        `bson:"views"`
	Data      ResourceData `bson:"data"`
}

type ResourceRepo interface {
	FindByID(ctx context.Context, id int64) (*Resource, error)
	ListTopByArea(context.Context, interface{}) (*[]Resource, error)
	FetchMapByIds(context.Context, []int64) (map[int64]*Resource, error)
	Search(context.Context, string) ([]*Resource, error)
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

func (uc *ResourceUsecase) Search(ctx context.Context, keyword string) ([]*Resource, error) {
	return uc.repo.Search(ctx, keyword)
}

func (uc *ResourceUsecase) FetchMapByIds(ctx context.Context, ids []int64) (map[int64]*Resource, error) {
	return uc.repo.FetchMapByIds(ctx, ids)
}

func (uc *ResourceUsecase) FindByID(ctx context.Context, id int64) (*Resource, error) {
	return uc.repo.FindByID(ctx, id)
}
