package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/common/model"
)

type RepositoryInterface interface {
	InsertUser(ctx context.Context, in model.InsertUserReq) (out model.User, err error)
	GetUserPassword(ctx context.Context, in model.GetUserPasswordReq) (out model.GetUserPasswordRes, err error)
	AddUserLoginCount(ctx context.Context, in model.AddUserLoginCountReq) (err error)
	GetUserByUserID(ctx context.Context, in model.GetUserByUserIDReq) (out model.User, err error)
	UpdateUserByUserID(ctx context.Context, in model.UpdateUserByUserIDReq) (err error)
}
