package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/SawitProRecruitment/UserService/common/model"
	"github.com/SawitProRecruitment/UserService/common/validator"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetProfile(ctx echo.Context) error {

	userID, ok := ctx.Get("user_id").(int64)
	if !ok {
		log.Printf("[handler.GetProfile] fail get user id")
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: "InternalServerError"})
	}

	user, err := s.Repository.GetUserByUserID(context.Background(), model.GetUserByUserIDReq{
		UserID: int64(userID),
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

	userID, ok := ctx.Get("user_id").(int64)
	if !ok {
		log.Printf("[handler.GetProfile] fail get user id")
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: "InternalServerError"})
	}
	var fullName, phoneNumber string
	if request.FullName != nil {
		fullName = *request.FullName
	}
	if request.PhoneNumber != nil {
		phoneNumber = *request.PhoneNumber
	}
	err = s.Repository.UpdateUserByUserID(context.Background(), model.UpdateUserByUserIDReq{
		UserID:      userID,
		PhoneNumber: phoneNumber,
		FullName:    fullName,
	})
	if err != nil {
		log.Printf("[handler.PutProfile] UpdateUserByUserID err:  %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, "success")
}
