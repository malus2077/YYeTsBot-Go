package service

import (
	pb "YYeTsBot-Go/api/yyets/v1"
	"YYeTsBot-Go/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
)

type UserService struct {
	pb.UnimplementedUserServer
	userUseCase    *biz.UserUseCase
	captchaUseCase *biz.CaptchaUsecase
}

func NewUserService(userUseCase *biz.UserUseCase, captchaUseCase *biz.CaptchaUsecase) *UserService {
	return &UserService{
		userUseCase:    userUseCase,
		captchaUseCase: captchaUseCase,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	passVerify, err := s.captchaUseCase.Verify(ctx, req.CaptchaId, req.Captcha)

	if err != nil {
		return nil, err
	}
	if !passVerify {
		return nil, errors.New("captcha verify failed")
	}

	user, err := s.userUseCase.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 生成token
	token, err := s.userUseCase.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	reply := &pb.CreateUserReply{
		Username:    user.UserName,
		Group:       user.Group,
		StatusCode:  200,
		Message:     "登录成功",
		AccessToken: token,
	}
	return reply, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	token, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, errors.New("token not found")
	}

	id, err := token.GetIssuer()
	user, err := s.userUseCase.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Username:  user.UserName,
		Date:      user.Date,
		Browser:   user.Browser,
		Hash:      user.Hash,
		LastDate:  user.LastDate,
		LastIp:    user.LastIP,
		Group:     user.Group,
		HasAvatar: user.Avatar != nil,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
