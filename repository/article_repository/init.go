package article_repository

import (
	"context"
	"database/sql"
	"github.com/agisnur24/teman-piknik.git/domain/entity"
)

type ArticleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, article *entity.Article) *entity.Article
	Update(ctx context.Context, tx *sql.Tx, article *entity.Article) *entity.Article
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*entity.Article, error)
	FindByTitle(ctx context.Context, tx *sql.Tx, title string) (*entity.Article, error)
	FindAll(ctx context.Context, tx *sql.Tx) []*entity.Article
}
