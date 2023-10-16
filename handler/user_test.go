package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/common/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestServer_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := repository.NewMockRepositoryInterface(ctrl)

	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		mock    func(ctx echo.Context) echo.Context
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("GET", "/profile", bytes.NewBufferString(
						``,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func(ctx echo.Context) echo.Context {
				ctx.Set("user_id", int64(1))
				mockRepository.EXPECT().GetUserByUserID(context.Background(), model.GetUserByUserIDReq{
					UserID: 1,
				}).Return(model.User{
					UserID:      1,
					FullName:    "a",
					PhoneNumber: "a",
					LoginCount:  0,
				}, nil).Times(1)

				return ctx
			},
		},
		{
			name: "fail get user id",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("GET", "/profile", bytes.NewBufferString(
						``,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func(ctx echo.Context) echo.Context {

				return ctx
			},
		},
		{
			name: "fail GetUserByUserID",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("GET", "/profile", bytes.NewBufferString(
						``,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func(ctx echo.Context) echo.Context {
				ctx.Set("user_id", int64(1))
				mockRepository.EXPECT().GetUserByUserID(context.Background(), model.GetUserByUserIDReq{
					UserID: 1,
				}).Return(model.User{}, errors.New("a")).Times(1)

				return ctx
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = tt.mock(tt.args.ctx)
			s := &Server{
				Repository: mockRepository,
			}

			if err := s.GetProfile(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_PutProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := repository.NewMockRepositoryInterface(ctrl)

	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		mock    func(ctx echo.Context) echo.Context
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("POST", "/profile", bytes.NewBufferString(
						`{"full_name":"test1","phone_number":"+6281123123"}`,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func(ctx echo.Context) echo.Context {
				ctx.Set("user_id", int64(1))
				mockRepository.EXPECT().UpdateUserByUserID(context.Background(), model.UpdateUserByUserIDReq{
					UserID:      1,
					PhoneNumber: "+6281123123",
					FullName:    "test1",
				}).Return(nil).Times(1)
				return ctx
			},
		},
		{
			name: "fail get user id",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("GET", "/profile", bytes.NewBufferString(
						``,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func(ctx echo.Context) echo.Context {

				return ctx
			},
		},
		{
			name: "fail UpdateUserByUserID",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("POST", "/profile", bytes.NewBufferString(
						`{"full_name":"test1","phone_number":"+6281123123"}`,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func(ctx echo.Context) echo.Context {
				ctx.Set("user_id", int64(1))
				mockRepository.EXPECT().UpdateUserByUserID(context.Background(), model.UpdateUserByUserIDReq{
					UserID:      1,
					PhoneNumber: "+6281123123",
					FullName:    "test1",
				}).Return(errors.New("a")).Times(1)
				return ctx
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = tt.mock(tt.args.ctx)
			s := &Server{
				Repository: mockRepository,
			}
			if err := s.PutProfile(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.PutProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
