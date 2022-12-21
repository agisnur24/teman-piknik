package article_service

import (
	"context"
	"github.com/agisnur24/teman-piknik.git/web/http_article"
)

type ArticleService interface {
	Save(ctx context.Context, request *http_article.ArticlePostRequest) *http_article.ArticleResponse
	Update(ctx context.Context, request *http_article.ArticleUpdateRequest) *http_article.ArticleResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) *http_article.ArticleResponse
	FindByTitle(ctx context.Context, title string) *http_article.ArticleResponse
	FindAll(ctx context.Context) []*http_article.ArticleResponse
}
