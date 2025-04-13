package service

import (
	"YYeTsBot-Go/internal/biz"
	"context"

	pb "YYeTsBot-Go/api/yyets/v1"
)

type CaptchaService struct {
	pb.UnimplementedCaptchaServer
	uc biz.CaptchaUsecase
}

func NewCaptchaService(uc *biz.CaptchaUsecase) *CaptchaService {
	return &CaptchaService{
		uc: *uc,
	}
}

func (s *CaptchaService) CreateCaptcha(ctx context.Context, req *pb.CreateCaptchaRequest) (*pb.CreateCaptchaReply, error) {
	return &pb.CreateCaptchaReply{}, nil
}
func (s *CaptchaService) UpdateCaptcha(ctx context.Context, req *pb.UpdateCaptchaRequest) (*pb.UpdateCaptchaReply, error) {
	return &pb.UpdateCaptchaReply{}, nil
}
func (s *CaptchaService) DeleteCaptcha(ctx context.Context, req *pb.DeleteCaptchaRequest) (*pb.DeleteCaptchaReply, error) {
	return &pb.DeleteCaptchaReply{}, nil
}
func (s *CaptchaService) GetCaptcha(ctx context.Context, req *pb.GetCaptchaRequest) (*pb.GetCaptchaReply, error) {
	captcha, err := s.uc.Generate(ctx, &biz.Captcha{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.GetCaptchaReply{
		Base64: captcha.Base64,
	}, nil
}
func (s *CaptchaService) ListCaptcha(ctx context.Context, req *pb.ListCaptchaRequest) (*pb.ListCaptchaReply, error) {
	return &pb.ListCaptchaReply{}, nil
}
