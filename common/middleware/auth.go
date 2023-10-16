package middleware

import (
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/common/jwt"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) AuthTokenValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if _, ok := m.ByPassAuthEndpoint[c.Request().Method+" "+c.Path()]; !ok {
			authorization := c.Request().Header.Get("Authorization")
			token := strings.TrimPrefix(authorization, "Bearer ")

			jwtResult, err := jwt.Validate(token)
			if err != nil {
				return c.JSON(http.StatusForbidden, generated.ErrorResponse{Message: err.Error()})
			}

			c.Set("user_id", jwtResult.UserID)
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
