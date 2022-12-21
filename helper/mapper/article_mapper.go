package mapper

import (
	"github.com/agisnur24/teman-piknik.git/domain/entity"
	"github.com/agisnur24/teman-piknik.git/web/http_article"
)

type ArticleMapper struct {
}

func (a ArticleMapper) FromDomainToResponse(domain *entity.Article) *http_article.ArticleResponse {
	return &http_article.ArticleResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		Image:       domain.Image,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (a ArticleMapper) FromDomainToMultiResponses(domain []*entity.Article) []*http_article.ArticleResponse {
	var articleResponses []*http_article.ArticleResponse
	for _, articleResponse := range domain {
		articleResponses = append(articleResponses, a.FromDomainToResponse(articleResponse))
	}
	return articleResponses
}
