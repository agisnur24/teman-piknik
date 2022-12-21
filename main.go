package main

import (
	"github.com/agisnur24/teman-piknik.git/app"
	"github.com/agisnur24/teman-piknik.git/app/db_connection"
	"github.com/agisnur24/teman-piknik.git/controller/article_controller"
	"github.com/agisnur24/teman-piknik.git/helper"
	"github.com/agisnur24/teman-piknik.git/middleware"
	"github.com/agisnur24/teman-piknik.git/repository/article_repository"
	"github.com/agisnur24/teman-piknik.git/service/article_service"
	"github.com/go-playground/validator"
	"net/http"
)

func main() {
	var (
		db       = db_connection.NewDB()
		validate = validator.New()
	)

	articleRepository := article_repository.NewArticleMysqlRepository()
	articleService := article_service.NewArticleService(articleRepository, db, validate)
	articleController := article_controller.NewArticleController(articleService)

	router := app.NewRouter(articleController)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
