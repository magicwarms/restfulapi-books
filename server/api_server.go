package server

import (
	"fmt"
	"net/http"
	"os"
	"restfulapi-books/apps/constants"
	"restfulapi-books/apps/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer(
	// accountHandler *accounts.AccountHandler,
	logger utils.Logger,
) *APIServer {
	echoApp := echo.New()

	echoApp.Debug = os.Getenv("APPLICATION_ENV") == constants.DEV_ENV

	echoApp.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogHost:   true,
		LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info(ctx,
				"API-REQUEST-LOG",
				utils.Fields{
					"status":    v.Status,
					"method":    v.Method,
					"uri":       v.URI,
					"userAgent": v.UserAgent,
					"remoteIP":  v.RemoteIP,
					"host":      v.Host,
				},
			)
			return nil
		},
	}))

	echoApp.Use(middleware.CORS())

	echoApp.Use(middleware.Gzip())

	echoApp.Use(middleware.Secure())

	echoApp.Use(middleware.RequestID())

	echoApp.Use(middleware.Recover())

	baseRoute := echoApp.Group("/api/v1")

	baseRoute.GET("/health", func(ctx echo.Context) error {
		return utils.AppResponse(ctx, http.StatusOK, "Backend Up and Running")
	})

	// -----------------ROUTES----------------//
	// accountRoute := baseRoute.Group("/accounts")
	// accountRoute.GET("/find", accountHandler.FindAccount)
	// accountRoute.POST("/transfer", accountHandler.ProcessMoneyTransferToAccount)
	// accountRoute.GET("/check", accountHandler.CheckTransactionStatus)

	return &APIServer{
		httpServer: &http.Server{
			Addr: "0.0.0.0:" + os.Getenv("PORT"),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      echoApp,
		},
	}
}

func (s *APIServer) Start() error {

	fmt.Println("⚡️ [" + os.Getenv("APPLICATION_ENV") + "] - " + os.Getenv("APP_NAME") + " IS RUNNING ON PORT - " + "http://localhost:" + os.Getenv("PORT"))

	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
