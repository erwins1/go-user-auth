package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/common/model"
)

func TestRepository_InsertUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	type args struct {
		ctx context.Context
		in  model.InsertUserReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantOut model.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.InsertUserReq{
					FullName:    "a",
					PhoneNumber: "a",
					Password:    "a",
					Salt:        "a",
				},
			},
			mock: func() {
				mock.ExpectQuery(queryInsertUser).WithArgs("a", "a", "a", "a").
					WillReturnRows(
						sqlmock.NewRows([]string{"user_id", "full_name", "phone_number"}).
							AddRow(1, "a", "a"),
					)
			},
			wantOut: model.User{
				UserID:      1,
				FullName:    "a",
				PhoneNumber: "a",
				LoginCount:  0,
			},
		},
		{
			name: "fail",
			args: args{
				ctx: context.Background(),
				in: model.InsertUserReq{
					FullName:    "a",
					PhoneNumber: "a",
					Password:    "a",
					Salt:        "a",
				},
			},
			mock: func() {
				mock.ExpectQuery(queryInsertUser).WithArgs("a", "a", "a", "a").
					WillReturnRows(
						sqlmock.NewRows([]string{"user_id", "full_name", "phone_number"}),
					)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: db,
			}
			gotOut, err := r.InsertUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Repository.InsertUser() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestRepository_GetUserPassword(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	type args struct {
		ctx context.Context
		in  model.GetUserPasswordReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantOut model.GetUserPasswordRes
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.GetUserPasswordReq{
					PhoneNumber: "a",
				},
			},
			mock: func() {
				mock.ExpectQuery(queryGetUserPassword).WithArgs("a").
					WillReturnRows(
						sqlmock.NewRows([]string{"user_id", "password", "salt"}).
							AddRow(1, "a", "a"),
					)
			},
			wantOut: model.GetUserPasswordRes{
				UserID:   1,
				Password: "a",
				Salt:     "a",
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.GetUserPasswordReq{
					PhoneNumber: "a",
				},
			},
			mock: func() {
				mock.ExpectQuery(queryGetUserPassword).WithArgs("a").
					WillReturnRows(
						sqlmock.NewRows([]string{"user_id", "password", "salt"}),
					)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: db,
			}
			gotOut, err := r.GetUserPassword(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetUserPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Repository.GetUserPassword() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestRepository_AddUserLoginCount(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	type args struct {
		ctx context.Context
		in  model.AddUserLoginCountReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.AddUserLoginCountReq{
					PhoneNumber: "a",
				},
			},
			mock: func() {
				mock.ExpectExec(queryUpdateUserLoginCount).WithArgs("a").WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "err",
			args: args{
				ctx: context.Background(),
				in: model.AddUserLoginCountReq{
					PhoneNumber: "a",
				},
			},
			mock: func() {
				mock.ExpectExec(queryUpdateUserLoginCount).WithArgs("a").WillReturnError(errors.New("a"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: db,
			}
			if err := r.AddUserLoginCount(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Repository.AddUserLoginCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetUserByUserID(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	type args struct {
		ctx context.Context
		in  model.GetUserByUserIDReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantOut model.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.GetUserByUserIDReq{
					UserID: 1,
				},
			},
			mock: func() {
				mock.ExpectQuery(queryGetUserByUserID).WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"user_id", "full_name", "phone_number", "login_count"}).
							AddRow(1, "a", "a", 1),
					)
			},
			wantOut: model.User{
				UserID:      1,
				FullName:    "a",
				PhoneNumber: "a",
				LoginCount:  1,
			},
		},
		{
			name: "err",
			args: args{
				ctx: context.Background(),
				in: model.GetUserByUserIDReq{
					UserID: 1,
				},
			},
			mock: func() {
				mock.ExpectQuery(queryGetUserByUserID).WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"user_id", "full_name", "phone_number", "login_count"}),
					)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: db,
			}
			gotOut, err := r.GetUserByUserID(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetUserByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Repository.GetUserByUserID() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestRepository_UpdateUserByUserID(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	type args struct {
		ctx context.Context
		in  model.UpdateUserByUserIDReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in: model.UpdateUserByUserIDReq{
					UserID:      1,
					PhoneNumber: "a",
					FullName:    "a",
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE users SET phone_number = $1 , full_name = $2 WHERE user_id = $3").WithArgs("a", "a", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "err",
			args: args{
				ctx: context.Background(),
				in: model.UpdateUserByUserIDReq{
					UserID:      1,
					PhoneNumber: "a",
					FullName:    "a",
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE users SET phone_number = $1 , full_name = $2 WHERE user_id = $3").WithArgs("a", "a", 1).WillReturnError(errors.New("a"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: db,
			}
			if err := r.UpdateUserByUserID(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Repository.UpdateUserByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
