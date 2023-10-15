package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/SawitProRecruitment/UserService/common/auth"
	"github.com/SawitProRecruitment/UserService/common/jwt"
	"github.com/SawitProRecruitment/UserService/common/model"
	"github.com/SawitProRecruitment/UserService/common/validator"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

// handler for register user
func (s *Server) PostRegister(ctx echo.Context) error {

	jsonBody := generated.PostRegisterJSONBody{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("[handler.PostRegister] json decode err:  %v", err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}

	err = validator.ValidateRegisterJSONBody(jsonBody)
	if err != nil {
		log.Printf("[handler.PostRegister] validator err: %v", err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	salt := auth.GenerateSalt()
	password, err := auth.HashPassword(*jsonBody.Password, salt)
	if err != nil {
		log.Printf("[handler.PostRegister] fail hashing, err : %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	user, err := s.Repository.InsertUser(context.Background(), model.InsertUserReq{
		FullName:    *jsonBody.FullName,
		PhoneNumber: *jsonBody.PhoneNumber,
		Password:    password,
		Salt:        salt,
	})
	if err != nil {
		log.Printf("[handler.PostRegister] fail insert user, err : %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, user.UserID)
}

func (s *Server) PostLogin(ctx echo.Context) error {
	request := generated.PostLoginJSONRequestBody{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		log.Printf("[handler.PostLogin] json decode err:  %v", err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	err = validator.ValidateLoginJSONBody(request)
	if err != nil {
		log.Printf("[handler.PostLogin] validate err:  %v", err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}

	user, err := s.Repository.GetUserPassword(context.Background(), model.GetUserPasswordReq{
		PhoneNumber: *request.PhoneNumber,
	})
	if err != nil {
		log.Printf("[handler.PostLogin] fail get user password, err : %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}

	err = auth.ValidatePassword(*request.Password, user.Password, user.Salt)
	if err != nil {
		log.Printf("[handler.PostLogin] fail ValidatePassword, err : %v", err)
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{Message: "wrong password"})
	}
	err = s.Repository.AddUserLoginCount(context.Background(), model.AddUserLoginCountReq{
		PhoneNumber: *request.PhoneNumber,
	})
	if err != nil {
		log.Printf("[handler.PostLogin] fail AddUserLoginCount, err : %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	token, err := jwt.GenerateJWTToken(user.UserID)
	if err != nil {
		log.Printf("[handler.PostLogin] fail GenerateJWTToken, err : %v", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
