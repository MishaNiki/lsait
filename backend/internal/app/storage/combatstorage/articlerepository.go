package combatstorage

import (
	"database/sql"

	"github.com/MishaNiki/lsait/backend/internal/app/model"
)

type ArticleRepository struct {
	storage *CombatStorage
}

type articleSQL struct {
	id          int
	title       sql.NullString
	description sql.NullString
	date        sql.NullString
	text        sql.NullString
}

func (ar *ArticleRepository) Create(a *model.Article) {
	ar.storage.pgdb.QueryRow(
		`UPDATE "shLSAIT"."objArticle"
			SET title=$1, description=$2, text=$3, id_theme=$4, article=true
			WHERE id=$5 AND id_auth=$6 ;`,
		a.Title,
		a.Description,
		a.Text,
		a.Theme,
		a.ID,
		a.Auth.ID,
	)
}

func (ar *ArticleRepository) CreateDraft(a *model.Article) error {
	if err := ar.storage.pgdb.QueryRow(
		`INSERT INTO "shLSAIT"."objArticle"(
			title, description, id_auth, article)
			VALUES ($1, $2, $3, false) RETURNING id;`,
		a.Title,
		a.Description,
		a.Auth.ID,
	).Scan(&a.ID); err != nil {
		return err
	}
	return nil
}

func (ar *ArticleRepository) Update(a *model.Article) {
	ar.storage.pgdb.QueryRow(
		`UPDATE "shLSAIT"."objArticle"
			SET title=$1, description=$2, text=$3, id_theme=$4
			WHERE id=$5 AND id_auth=$6 ;`,
		a.Title,
		a.Description,
		a.Text,
		a.Theme,
		a.ID,
		a.Auth.ID,
	)
}

func (ar *ArticleRepository) GetByID(id int) (*model.Article, error) {
	a := &model.Article{ID: id}
	asql := &articleSQL{}
	if err := ar.storage.pgdb.QueryRow(
		`SELECT title, description, text, date
			FROM "shLSAIT"."objArticle"
			WHERE id=$1 AND article=true;`,
		id,
	).Scan(
		&asql.title,
		&asql.description,
		&asql.text,
		&a.Date,
	); err != nil {
		return nil, err
	}
	a.Title = asql.title.String
	a.Description = asql.description.String
	a.Text = asql.text.String
	return a, nil
}

func (ar *ArticleRepository) GetEditByID(id, idUser int) (*model.Article, error) {
	a := &model.Article{ID: id}
	asql := &articleSQL{}
	if err := ar.storage.pgdb.QueryRow(
		`SELECT title, description, text, date, article
			FROM "shLSAIT"."objArticle"
			WHERE id=$1 AND id_auth=$2;`,
		id,
		idUser,
	).Scan(
		&asql.title,
		&asql.description,
		&asql.text,
		&a.Date,
		&a.Article,
	); err != nil {
		return nil, err
	}
	a.Title = asql.title.String
	a.Description = asql.description.String
	a.Text = asql.text.String
	return a, nil
}

func (ar *ArticleRepository) DeleteByID(id int, idUser int) {
	ar.storage.pgdb.QueryRow(
		`DELETE FROM "shLSAIT"."objArticle"
			WHERE id=$1 AND id_auth=$2;`,
		id,
		idUser,
	)
}
