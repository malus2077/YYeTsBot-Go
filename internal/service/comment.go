package service

import (
	"YYeTsBot-Go/internal/biz"
	"context"

	pb "YYeTsBot-Go/api/yyets/v1"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	commentUseCase biz.CommentUsecase
	userUseCase    biz.UserUseCase
}

func NewCommentService(commentUseCase *biz.CommentUsecase, userUseCase *biz.UserUseCase) *CommentService {
	return &CommentService{commentUseCase: *commentUseCase, userUseCase: *userUseCase}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	return &pb.CreateCommentReply{}, nil
}
func (s *CommentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentReply, error) {
	return &pb.UpdateCommentReply{}, nil
}
func (s *CommentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	return &pb.DeleteCommentReply{}, nil
}
func (s *CommentService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	return &pb.GetCommentReply{}, nil
}
func (s *CommentService) ListComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	reply := &pb.ListCommentReply{}

	comments, count, err := s.commentUseCase.List(ctx, req.ResourceId, req.Page, req.Size, req.Sort)
	if err != nil {
		return nil, err
	}

	//todo: comment domain 聚合根
	userNames := make([]string, 0)
	for _, comment := range comments {
		userNames = append(userNames, comment.Username)
	}
	userDataMap, err := s.userUseCase.UserGroup(ctx, userNames)

	for _, comment := range comments {
		user := userDataMap[comment.Username]
		group := []string{"user"}
		if len(user.Group) > 0 {
			group = user.Group
		}

		item := &pb.CommentItem{
			Browser:   comment.Browser,
			Content:   comment.Content,
			Date:      comment.Date,
			Id:        comment.ID.Hex(),
			Group:     group,
			HasAvatar: user.Avatar != nil,
		}
		reply.Data = append(reply.Data, item)
	}

	reply.Count = count
	reply.ResourceId = req.ResourceId
	return reply, nil
}
func (s *CommentService) ListLatestComment(ctx context.Context, req *pb.ListLatestCommentRequest) (*pb.ListLatestCommentReply, error) {
	reply := &pb.ListLatestCommentReply{}

	userNames := make([]string, 0)
	comments, err := s.commentUseCase.ListLatest(ctx, req.Size)

	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		userNames = append(userNames, comment.Username)
	}
	userDataMap, err := s.userUseCase.UserGroup(ctx, userNames)

	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		user := userDataMap[comment.Username]
		group := []string{"user"}
		if len(user.Group) > 0 {
			group = user.Group
		}

		item := &pb.CommentItem{
			Browser:   comment.Browser,
			Content:   comment.Content,
			Date:      comment.Date,
			Id:        comment.ID.Hex(),
			Group:     group,
			HasAvatar: user.Avatar != nil,
		}

		reply.Data = append(reply.Data, item)
	}

	reply.Count, err = s.commentUseCase.Count(ctx)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
