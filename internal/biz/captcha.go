package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Captcha struct {
	Id     string
	Base64 string
}

type CaptchaRepo interface {
	Generate(ctx context.Context, requestId string) (*Captcha, error)
	Verify(ctx context.Context, requestId, userInput string) (bool, error)
}
type CaptchaUsecase struct {
	repo CaptchaRepo
	log  *log.Helper
}

func NewCaptchaUsecase(repo CaptchaRepo, logger log.Logger) *CaptchaUsecase {
	return &CaptchaUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CaptchaUsecase) Generate(ctx context.Context, captcha *Captcha) (*Captcha, error) {
	return uc.repo.Generate(ctx, captcha.Id)
}

func (uc *CaptchaUsecase) Verify(ctx context.Context, captchaId, userInput string) (bool, error) {
	return uc.repo.Verify(ctx, captchaId, userInput)
}
