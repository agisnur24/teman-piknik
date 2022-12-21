package app

import (
	"github.com/agisnur24/teman-piknik.git/controller/article_controller"
	"github.com/agisnur24/teman-piknik.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(article article_controller.ArticleController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/articles", article.Post)
	router.PUT("/api/articles/:id", article.Update)
	router.DELETE("/api/articles/:id", article.Delete)
	router.GET("/api/articles/:id", article.FindById)
	router.GET("/api/article/:title", article.FindByTitle)
	router.GET("/api/articles", article.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
