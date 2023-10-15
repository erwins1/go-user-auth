package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func TestServer_PostRegister(t *testing.T) {
	mockRepository := new(repository.MockRepositoryInterface)

	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{

		{
			name: "fail validate",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("POST", "/register", bytes.NewBufferString(
						`{"full_name":"test test","phone_number":"+6281123123","password":"Password123"}`,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			s := &Server{
				Repository: mockRepository,
			}
			if err := s.PostRegister(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.PostRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_PostLogin(t *testing.T) {
	mockRepository := new(repository.MockRepositoryInterface)

	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "fail validate",
			args: args{
				ctx: echo.New().NewContext(
					httptest.NewRequest("POST", "/login", bytes.NewBufferString(
						`{"phone_number":"+6281123123"}`,
					)),
					httptest.NewRecorder(),
				),
			},
			mock: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			s := &Server{
				Repository: mockRepository,
			}
			if err := s.PostLogin(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.PostLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
