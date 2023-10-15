package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/common/jwt"
	"github.com/SawitProRecruitment/UserService/common/model"
	"github.com/SawitProRecruitment/UserService/common/validator"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetProfile(ctx echo.Context) error {

	authorization := ctx.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authorization, "Bearer ")

	jwtResult, err := jwt.Validate(token)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{Message: err.Error()})
	}
	user, err := s.Repository.GetUserByUserID(context.Background(), model.GetUserByUserIDReq{
		UserID: jwtResult.UserID,
	})
	if err != nil {
		log.Printf("[handler.GetProfile] GetUserByUserID err:  %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}

func (s *Server) PutProfile(ctx echo.Context) error {

	request := generated.PutProfileJSONRequestBody{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		log.Printf("[handler.PutProfile] json decode err:  %v", err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	err = validator.ValidateBodyUpdateUser(request)
	if err != nil {
		log.Printf("[handler.PutProfile] validate err:  %v", err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}

	authorization := ctx.Request().Header.Get("Authorization")
	token := strings.TrimPrefix(authorization, "Bearer ")

	jwtResult, err := jwt.Validate(token)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{Message: err.Error()})
	}
	err = s.Repository.UpdateUserByUserID(context.Background(), model.UpdateUserByUserIDReq{
		UserID:      jwtResult.UserID,
		PhoneNumber: *request.PhoneNumber,
		FullName:    *request.FullName,
	})
	if err != nil {
		log.Printf("[handler.PutProfile] UpdateUserByUserID err:  %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, "success")
}
