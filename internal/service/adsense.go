package service

import (
	"context"

	pb "YYeTsBot-Go/api/yyets/v1"
)

type AdsenseService struct {
	pb.UnimplementedAdsenseServer
}

func NewAdsenseService() *AdsenseService {
	return &AdsenseService{}
}

func (s *AdsenseService) CreateAdsense(ctx context.Context, req *pb.CreateAdsenseRequest) (*pb.CreateAdsenseReply, error) {
	return &pb.CreateAdsenseReply{}, nil
}
func (s *AdsenseService) UpdateAdsense(ctx context.Context, req *pb.UpdateAdsenseRequest) (*pb.UpdateAdsenseReply, error) {
	return &pb.UpdateAdsenseReply{}, nil
}
func (s *AdsenseService) DeleteAdsense(ctx context.Context, req *pb.DeleteAdsenseRequest) (*pb.DeleteAdsenseReply, error) {
	return &pb.DeleteAdsenseReply{}, nil
}
func (s *AdsenseService) GetAdsense(ctx context.Context, req *pb.GetAdsenseRequest) (*pb.GetAdsenseReply, error) {
	return &pb.GetAdsenseReply{}, nil
}
func (s *AdsenseService) ListAdsense(ctx context.Context, req *pb.ListAdsenseRequest) (*pb.ListAdsenseReply, error) {
	return &pb.ListAdsenseReply{
		Data: make([]string, 1),
	}, nil
}
