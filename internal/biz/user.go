package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
)

type User struct {
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetStoreClient(ctx context.Context) *minio.Client
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) GetStoreClient(ctx context.Context) *minio.Client {
	return uc.ur.GetStoreClient(ctx)
}
