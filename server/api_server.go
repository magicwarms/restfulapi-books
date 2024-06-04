package server

import (
	"fmt"
	"net/http"
	"os"
	"restfulapi-books/apps/authors"
	"restfulapi-books/apps/books"
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
	bookHandler *books.BookHandler,
	authorHandler *authors.AuthorHandler,
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
	bookRoute := baseRoute.Group("/books")
	bookRoute.GET("/get", bookHandler.FindBook)
	bookRoute.POST("/store", bookHandler.AddBook)
	bookRoute.GET("/all", bookHandler.GetAllBooks)
	bookRoute.PUT("/update", bookHandler.UpdateBook)
	bookRoute.DELETE("/delete", bookHandler.DeleteBook)

	authorRoute := baseRoute.Group("/authors")
	authorRoute.GET("/get", authorHandler.FindAuthor)
	authorRoute.POST("/store", authorHandler.AddAuthor)
	authorRoute.GET("/all", authorHandler.GetAllAuthors)
	authorRoute.PUT("/update", authorHandler.UpdateAuthor)
	authorRoute.DELETE("/delete", authorHandler.DeleteAuthor)

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
