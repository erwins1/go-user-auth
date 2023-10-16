package main

import (
	"os"

	"github.com/SawitProRecruitment/UserService/common/middleware"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()
	initMiddleware(e)
	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}

func initMiddleware(e *echo.Echo) {
	m := middleware.Init(middleware.Middleware{
		ByPassAuthEndpoint: map[string]bool{
			"POST /register": true,
			"POST /login":    true,
		},
	})
	e.Use(m.AuthTokenValidation)
}
