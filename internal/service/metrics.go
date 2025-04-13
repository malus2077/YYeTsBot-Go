package service

import (
	"context"

	pb "YYeTsBot-Go/api/yyets/v1"
)

type MetricsService struct {
	pb.UnimplementedMetricsServer
}

func NewMetricsService() *MetricsService {
	return &MetricsService{}
}

func (s *MetricsService) CreateMetrics(ctx context.Context, req *pb.CreateMetricsRequest) (*pb.CreateMetricsReply, error) {
	return &pb.CreateMetricsReply{}, nil
}
func (s *MetricsService) UpdateMetrics(ctx context.Context, req *pb.UpdateMetricsRequest) (*pb.UpdateMetricsReply, error) {
	return &pb.UpdateMetricsReply{}, nil
}
func (s *MetricsService) DeleteMetrics(ctx context.Context, req *pb.DeleteMetricsRequest) (*pb.DeleteMetricsReply, error) {
	return &pb.DeleteMetricsReply{}, nil
}
func (s *MetricsService) GetMetrics(ctx context.Context, req *pb.GetMetricsRequest) (*pb.GetMetricsReply, error) {
	return &pb.GetMetricsReply{}, nil
}
func (s *MetricsService) ListMetrics(ctx context.Context, req *pb.ListMetricsRequest) (*pb.ListMetricsReply, error) {
	return &pb.ListMetricsReply{}, nil
}
