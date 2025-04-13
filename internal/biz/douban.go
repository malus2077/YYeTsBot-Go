package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DoubanRaw struct {
	SearchUrl  string `bson:"search_url"`
	DetailUrl  string `bson:"detail_url"`
	SearchHtml string `bson:"search_html"`
	DetailHtml string `bson:"detail_html"`
}

type Douban struct {
	Id              bson.ObjectID `bson:"_id"`
	Name            string        `bson:"name" json:"name"`
	Raw             *DoubanRaw    `bson:"raw" json:"raw"`
	DoubanId        int64         `bson:"doubanId" json:"doubanId"`
	DoubanLink      string        `bson:"doubanLink" json:"doubanLink"`
	PosterLink      string        `bson:"posterLink" json:"posterLink"`
	PosterData      []byte        `bson:"posterData" json:"posterData"`
	ResourceId      int64         `bson:"resourceId" json:"resourceId"`
	Rating          string        `bson:"rating" json:"rating"`
	ReleaseDate     string        `bson:"releaseDate" json:"releaseDate"`
	Actors          []string      `bson:"actors" json:"actors"`
	Directors       []string      `bson:"directors" json:"directors"`
	Genres          []string      `bson:"genres" json:"genres"`
	Year            string        `bson:"year" json:"year"`
	EpisodeCount    string        `bson:"episodeCount" json:"episodeCount"`
	EpisodeDuration string        `bson:"episodeDuration" json:"episodeDuration"`
	Writers         []string      `bson:"writers" json:"writers"`
	Introduction    string        `bson:"introduction" json:"introduction"`
}
type DoubanRepo interface {
	FindByResourceId(ctx context.Context, resourceId int64) (*Douban, error)
	Insert(ctx context.Context, resourceName string, resourceId int64) (douban *Douban, err error)
}

type DoubanUsecase struct {
	doubanRepo   DoubanRepo
	resourceRepo ResourceRepo
	log          *log.Helper
}

func NewDoubanUsecase(doubanRepo DoubanRepo, resourceRepo ResourceRepo, logger log.Logger) *DoubanUsecase {
	return &DoubanUsecase{doubanRepo: doubanRepo, resourceRepo: resourceRepo, log: log.NewHelper(logger)}
}
func (uc *DoubanUsecase) FetchByResourceId(ctx context.Context, resourceId int64) (*Douban, error) {
	result, err := uc.doubanRepo.FindByResourceId(ctx, resourceId)
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result, nil
	}

	resource, err := uc.resourceRepo.FindByID(ctx, resourceId)
	if err != nil {
		return nil, err
	}

	result, err = uc.doubanRepo.Insert(ctx, resource.CnName, resourceId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
