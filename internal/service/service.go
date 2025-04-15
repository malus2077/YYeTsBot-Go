package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewResourceService, NewCommentService, NewAnnouncementService, NewMetricsService, NewAdsenseService, NewDoubanService, NewCaptchaService, NewUserService, NewNotificationService)
