package service

import (
	"YYeTsBot-Go/internal/biz"
	"context"

	pb "YYeTsBot-Go/api/yyets/v1"
)

type DoubanService struct {
	pb.UnimplementedDoubanServer
	uc *biz.DoubanUsecase
}

func NewDoubanService(uc *biz.DoubanUsecase) *DoubanService {
	return &DoubanService{
		uc: uc,
	}
}

func (s *DoubanService) CreateDouban(ctx context.Context, req *pb.CreateDoubanRequest) (*pb.CreateDoubanReply, error) {
	return &pb.CreateDoubanReply{}, nil
}
func (s *DoubanService) UpdateDouban(ctx context.Context, req *pb.UpdateDoubanRequest) (*pb.UpdateDoubanReply, error) {
	return &pb.UpdateDoubanReply{}, nil
}
func (s *DoubanService) DeleteDouban(ctx context.Context, req *pb.DeleteDoubanRequest) (*pb.DeleteDoubanReply, error) {
	return &pb.DeleteDoubanReply{}, nil
}
func (s *DoubanService) GetDouban(ctx context.Context, req *pb.GetDoubanRequest) (*pb.GetDoubanReply, error) {
	douban, err := s.uc.FetchByResourceId(ctx, req.ResourceId)
	if err != nil {
		return nil, err
	}
	reply := &pb.GetDoubanReply{
		Name:            douban.Name,
		DoubanId:        douban.DoubanId,
		DoubanLink:      douban.DoubanLink,
		PosterLink:      douban.PosterLink,
		ResourceId:      douban.ResourceId,
		Rating:          douban.Rating,
		Actors:          douban.Actors,
		Directors:       douban.Directors,
		Genre:           douban.Genres,
		ReleaseDate:     douban.ReleaseDate,
		EpisodeCount:    douban.EpisodeCount,
		EpisodeDuration: douban.EpisodeDuration,
		Writers:         douban.Writers,
		Year:            douban.Year,
		Introduction:    douban.Introduction,
	}

	return reply, nil
}
func (s *DoubanService) ListDouban(ctx context.Context, req *pb.ListDoubanRequest) (*pb.ListDoubanReply, error) {
	return &pb.ListDoubanReply{}, nil
}
