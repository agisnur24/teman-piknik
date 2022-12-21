package article_controller

import (
	"errors"
	"github.com/agisnur24/teman-piknik.git/helper"
	"github.com/agisnur24/teman-piknik.git/service/article_service"
	"github.com/agisnur24/teman-piknik.git/web"
	"github.com/agisnur24/teman-piknik.git/web/http_article"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ArticleControllerImplement struct {
	articleService article_service.ArticleService
}

func NewArticleController(svc article_service.ArticleService) ArticleController {
	return &ArticleControllerImplement{articleService: svc}
}

func (c *ArticleControllerImplement) Post(writer http.ResponseWriter, request *http.Request,
	params httprouter.Params) {

	articleRequest := &http_article.ArticlePostRequest{}
	helper.ReadFromRequestBody(request, articleRequest)

	articleResponse := c.articleService.Save(request.Context(), articleRequest)
	httpResponse := &web.HttpResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}

func (c *ArticleControllerImplement) Update(writer http.ResponseWriter,
	request *http.Request, params httprouter.Params) {

	updateRequest := http_article.ArticleUpdateRequest{}
	helper.ReadFromRequestBody(request, &updateRequest)

	articleId := params.ByName("id")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	updateRequest.ID = id

	articleResponse := c.articleService.Update(request.Context(), &updateRequest)
	webResponse := web.HttpResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *ArticleControllerImplement) Delete(writer http.ResponseWriter, request *http.Request,
	params httprouter.Params) {

	articleId := params.ByName("id")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	response := c.articleService.FindById(request.Context(), id)
	c.articleService.Delete(request.Context(), id)
	webResponse := web.HttpResponse{
		Code:   200,
		Status: "OK",
		Data:   "Artikel dengan judul " + response.Title + " telah terhapus",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *ArticleControllerImplement) FindById(writer http.ResponseWriter, request *http.Request,
	params httprouter.Params) {

	articleId := params.ByName("id")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	articleResponse := c.articleService.FindById(request.Context(), id)
	webResponse := web.HttpResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *ArticleControllerImplement) FindByTitle(writer http.ResponseWriter, request *http.Request,
	params httprouter.Params) {

	articleTitle := params.ByName("title")
	if articleTitle == "" {
		helper.PanicIfError(errors.New("Endpoint param tidak sesuai"))
	}

	articleResponse := c.articleService.FindByTitle(request.Context(), articleTitle)
	webResponse := web.HttpResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *ArticleControllerImplement) FindAll(writer http.ResponseWriter, request *http.Request,
	params httprouter.Params) {

	articleResponses := c.articleService.FindAll(request.Context())
	webResponse := &web.HttpResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
