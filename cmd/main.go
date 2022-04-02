package main

import (
	"bareksaIntern/src/news"
	"bareksaIntern/src/tags"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodPost, http.MethodOptions, http.MethodGet},
		AllowHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "CUSTOM-ALLOWED-HEADER", "Authorization", "Access-Key", "X-CSRF-Token,"},
	}))

	e.POST("/news/search_one", news.SearchOneNews())
	e.POST("/news/search_many", news.SearchManyNews())
	e.POST("/news/create", news.CreateNews())
	e.POST("/news/update", news.UpdateNews())
	e.POST("/news/delete", news.DeleteNews())
	e.POST("/news/filter_topic", news.FilterNewsByTopic())

	e.POST("/tags/create", tags.CreateTag())
	e.POST("/tags/search_one", tags.SearchOneTag())
	e.POST("/tags/update", tags.UpdateTag())
	e.POST("/tags/delete", tags.DeleteTag())

	e.Logger.Fatal(e.Start(":8080"))
}
