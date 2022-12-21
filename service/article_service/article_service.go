package article_service

import (
	"context"
	"database/sql"
	"github.com/agisnur24/teman-piknik.git/domain/entity"
	"github.com/agisnur24/teman-piknik.git/exception"
	"github.com/agisnur24/teman-piknik.git/helper"
	"github.com/agisnur24/teman-piknik.git/helper/mapper"
	"github.com/agisnur24/teman-piknik.git/repository/article_repository"
	"github.com/agisnur24/teman-piknik.git/web/http_article"
	"github.com/go-playground/validator"
	"time"
)

type ArticleServiceImplement struct {
	articleRepository article_repository.ArticleRepository
	db                *sql.DB
	validate          *validator.Validate
}

func NewArticleService(repo article_repository.ArticleRepository,
	db *sql.DB, validate *validator.Validate) ArticleService {
	return &ArticleServiceImplement{
		articleRepository: repo,
		db:                db,
		validate:          validate,
	}
}

func (s *ArticleServiceImplement) Save(ctx context.Context,
	request *http_article.ArticlePostRequest) *http_article.ArticleResponse {

	err := s.validate.Struct(request)
	helper.PanicIfError(err)

	tx, errTx := s.db.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	article := &entity.Article{
		Title:       request.Title,
		Description: request.Description,
		Image:       request.Image,
		CreatedAt:   time.Now().Format("02-01-2006 15:04:05"),
		UpdatedAt:   time.Now().Format("02-01-2006 15:04:05"),
	}

	article = s.articleRepository.Save(ctx, tx, article)
	return mapper.ArticleMapper{}.FromDomainToResponse(article)
}

func (s *ArticleServiceImplement) Update(ctx context.Context,
	request *http_article.ArticleUpdateRequest) *http_article.ArticleResponse {

	err := s.validate.Struct(request)
	helper.PanicIfError(err)

	tx, errTx := s.db.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	article, errRepo := s.articleRepository.FindById(ctx, tx, request.ID)
	if errRepo != nil {
		panic(exception.NewNotFoundError(errRepo.Error()))
	}

	article.Title = request.Title
	article.Description = request.Description
	article.Image = request.Image
	article.UpdatedAt = time.Now().Format("02-01-2006 15:04:05")

	article = s.articleRepository.Update(ctx, tx, article)

	return mapper.ArticleMapper{}.FromDomainToResponse(article)
}

func (s *ArticleServiceImplement) Delete(ctx context.Context, id int) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, errRepo := s.articleRepository.FindById(ctx, tx, id)
	if errRepo != nil {
		panic(exception.NewNotFoundError(errRepo.Error()))
	}

	s.articleRepository.Delete(ctx, tx, article.ID)
}

func (s *ArticleServiceImplement) FindById(ctx context.Context, id int) *http_article.ArticleResponse {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, errRepo := s.articleRepository.FindById(ctx, tx, id)
	if errRepo != nil {
		panic(exception.NewNotFoundError(errRepo.Error()))
	}

	return mapper.ArticleMapper{}.FromDomainToResponse(article)
}

func (s *ArticleServiceImplement) FindByTitle(ctx context.Context, title string) *http_article.ArticleResponse {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, errRepo := s.articleRepository.FindByTitle(ctx, tx, title)
	if errRepo != nil {
		panic(exception.NewNotFoundError(errRepo.Error()))
	}

	return mapper.ArticleMapper{}.FromDomainToResponse(article)
}

func (s *ArticleServiceImplement) FindAll(ctx context.Context) []*http_article.ArticleResponse {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	articles := s.articleRepository.FindAll(ctx, tx)

	return mapper.ArticleMapper{}.FromDomainToMultiResponses(articles)
}
