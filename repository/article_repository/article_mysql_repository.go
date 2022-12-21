package article_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agisnur24/teman-piknik.git/domain/entity"
	"github.com/agisnur24/teman-piknik.git/helper"
)

type ArticleMysqlRepository struct {
}

func NewArticleMysqlRepository() ArticleRepository {
	return &ArticleMysqlRepository{}
}

func (ArticleMysqlRepository) Save(ctx context.Context, tx *sql.Tx, article *entity.Article) *entity.Article {
	SQL := "INSERT INTO articles (title, description, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, article.Title, article.Description, article.Image,
		article.CreatedAt, article.UpdatedAt)
	helper.PanicIfError(err)

	id, errLastInsertId := result.LastInsertId()
	helper.PanicIfError(errLastInsertId)

	article.ID = int(id)

	return article
}

func (ArticleMysqlRepository) Update(ctx context.Context, tx *sql.Tx, article *entity.Article) *entity.Article {
	SQL := "UPDATE articles SET title = ?, description = ?, image = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, article.Title, article.Description, article.Image, article.UpdatedAt, article.ID)
	helper.PanicIfError(err)

	return article
}

func (ArticleMysqlRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM articles WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (ArticleMysqlRepository) FindById(ctx context.Context, tx *sql.Tx, id int) (*entity.Article, error) {
	SQL := "SELECT id, title, description, image, created_at, updated_at FROM articles WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	article := entity.Article{}
	if rows.Next() {
		errScan := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Image, &article.CreatedAt,
			&article.UpdatedAt)
		helper.PanicIfError(errScan)
		return &article, nil
	} else {
		return &article, errors.New("Artikel tidak ditemukan")
	}
}

func (ArticleMysqlRepository) FindByTitle(ctx context.Context, tx *sql.Tx, title string) (*entity.Article, error) {
	SQL := "SELECT id, title, description, image, created_at, updated_at FROM articles WHERE title = ?"
	rows, err := tx.QueryContext(ctx, SQL, title)
	helper.PanicIfError(err)
	defer rows.Close()

	article := entity.Article{}
	if rows.Next() {
		errScan := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Image, &article.CreatedAt,
			&article.UpdatedAt)
		helper.PanicIfError(errScan)
		return &article, nil
	} else {
		return &article, errors.New("Artikel tidak ditemukan")
	}
}

func (ArticleMysqlRepository) FindAll(ctx context.Context, tx *sql.Tx) []*entity.Article {
	SQL := "SELECT id, title, description, image, created_at, updated_at FROM articles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var articles []*entity.Article
	for rows.Next() {
		article := entity.Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Image, &article.CreatedAt,
			&article.UpdatedAt)
		helper.PanicIfError(err)
		articles = append(articles, &article)
	}
	return articles
}
