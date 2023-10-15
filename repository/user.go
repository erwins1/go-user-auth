package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/SawitProRecruitment/UserService/common/model"
)

func (r *Repository) InsertUser(ctx context.Context, in model.InsertUserReq) (out model.User, err error) {
	log.Println(" TEST ", in.Salt, " - ", in.Password)
	err = r.Db.QueryRowContext(ctx, queryInsertUser, in.FullName, in.PhoneNumber, in.Password, in.Salt).
		Scan(&out.UserID, &out.FullName, &out.PhoneNumber)
	if err != nil {
		log.Printf("[repository.InsertUser] fail create user, err: %v", err)
		return
	}
	return
}

func (r *Repository) GetUserPassword(ctx context.Context, in model.GetUserPasswordReq) (out model.GetUserPasswordRes, err error) {

	row := r.Db.QueryRowContext(ctx, queryGetUserPassword, in.PhoneNumber)

	err = row.Scan(&out.UserID, &out.Password, &out.Salt)
	if err != nil {
		log.Printf("[repository.GetUserPassword] fail get user password, err: %v", err)
		return out, err
	}

	return out, nil
}

func (r *Repository) AddUserLoginCount(ctx context.Context, in model.AddUserLoginCountReq) (err error) {

	_, err = r.Db.ExecContext(ctx, queryUpdateUserLoginCount, in.PhoneNumber)
	if err != nil {
		log.Printf("[repository.AddUserLoginCount] fail add login count, err: %v", err)
		return err
	}

	return nil
}

func (r *Repository) GetUserByUserID(ctx context.Context, in model.GetUserByUserIDReq) (out model.User, err error) {

	err = r.Db.QueryRowContext(ctx, queryGetUserByUserID, in.UserID).Scan(&out.UserID, &out.FullName, &out.PhoneNumber, &out.LoginCount)
	if err != nil {
		log.Printf("[repository.GetUserByUserID] fail get user profile, err: %v", err)
		return out, err
	}

	return out, err
}

func (r *Repository) UpdateUserByUserID(ctx context.Context, in model.UpdateUserByUserIDReq) (err error) {

	var (
		set []string
	)
	query := "UPDATE users"
	args := []interface{}{}
	if in.PhoneNumber != "" {
		args = append(args, in.PhoneNumber)
		set = append(set, fmt.Sprintf("phone_number = $%d", len(args)))
	}
	if in.FullName != "" {
		args = append(args, in.FullName)
		set = append(set, fmt.Sprintf("full_name = $%d", len(args)))
	}

	if len(set) > 0 {
		query += ` SET ` + strings.Join(set, ` , `)
	}

	args = append(args, in.UserID)
	query += fmt.Sprintf(" WHERE user_id = $%d", len(args))

	log.Println(query)

	_, err = r.Db.Exec(query, args...)
	return err
}
