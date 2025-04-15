package service

import (
	"context"

	pb "YYeTsBot-Go/api/yyets/v1"
)

type NotificationService struct {
	pb.UnimplementedNotificationServer
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationReply, error) {
	return &pb.CreateNotificationReply{}, nil
}
func (s *NotificationService) UpdateNotification(ctx context.Context, req *pb.UpdateNotificationRequest) (*pb.UpdateNotificationReply, error) {
	return &pb.UpdateNotificationReply{}, nil
}
func (s *NotificationService) DeleteNotification(ctx context.Context, req *pb.DeleteNotificationRequest) (*pb.DeleteNotificationReply, error) {
	return &pb.DeleteNotificationReply{}, nil
}
func (s *NotificationService) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationReply, error) {
	return &pb.GetNotificationReply{
		ReadCount:   0,
		UnreadCount: 0,
		Username:    "",
		ReadItem:    make([]*pb.NotificationItem, 0),
		UnreadItem:  make([]*pb.NotificationItem, 0),
	}, nil
}
func (s *NotificationService) ListNotification(ctx context.Context, req *pb.ListNotificationRequest) (*pb.ListNotificationReply, error) {
	return &pb.ListNotificationReply{}, nil
}
