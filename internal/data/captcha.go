package data

import (
	"YYeTsBot-Go/internal/biz"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

const (
	captchaCacheKey = "captcha:%s"
)

type captchaRepo struct {
	data *Data
	log  *log.Helper
}

func (c captchaRepo) Generate(ctx context.Context, requestId string) (*biz.Captcha, error) {
	cacheKey := fmt.Sprintf(captchaCacheKey, requestId)
	captchaID, err := c.data.cacheManager.Get(ctx, cacheKey)

	if err == nil && len(captchaID) > 0 {
		captcha.Reload(captchaID)
	} else {
		captchaID = captcha.NewLen(4)
	}

	var imgBuf bytes.Buffer
	if err := captcha.WriteImage(&imgBuf, captchaID, 240, 80); err != nil {
		return nil, err
	}

	base64Captcha := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	err = c.data.cacheManager.Set(ctx, cacheKey, captchaID, store.WithExpiration(5*time.Minute))
	if err != nil {
		return nil, err
	}
	return &biz.Captcha{
		Id:     captchaID,
		Base64: base64Captcha,
	}, nil
}

func (c captchaRepo) Verify(ctx context.Context, requestId, userInput string) (bool, error) {
	cacheKey := fmt.Sprintf(captchaCacheKey, requestId)
	captchaID, err := c.data.cacheManager.Get(ctx, cacheKey)
	defer c.data.cacheManager.Delete(ctx, cacheKey)

	if err != nil {
		return false, err
	}
	return captcha.VerifyString(captchaID, userInput), nil
}

func NewCaptchaRepo(data *Data, logger log.Logger) biz.CaptchaRepo {
	return &captchaRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
