package service

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

func NewRealWorldService(uc *biz.UserUsecase, logger log.Logger) *RealWorldService {
	ctx := context.Background()
	ossclient := uc.GetStoreClient(ctx)

	return &RealWorldService{uc: uc, oss: ossclient, log: log.NewHelper(logger)}
}
