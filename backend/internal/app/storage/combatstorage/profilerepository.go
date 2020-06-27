package combatstorage

import (
	"database/sql"
	"log"

	"github.com/MishaNiki/lsait/backend/internal/app/model"
)

type profileSQL struct {
	id          int
	name        sql.NullString
	surname     sql.NullString
	description sql.NullString
	position    sql.NullString
}

type ProfileRepository struct {
	storage *CombatStorage
}

// Create ...
func (pr *ProfileRepository) Create(*model.Profile) error {

	return nil
}

func (pr *ProfileRepository) Update(p *model.Profile) {
	pr.storage.pgdb.QueryRow(
		`UPDATE "shLSAIT"."objProfile"
			SET name=$1, surname=$2, "position"=$3, description=$4
			WHERE id = $5;`,
		p.Name,
		p.Surname,
		p.Position,
		p.Description,
		p.ID,
	)
}

func (pr *ProfileRepository) GetByID(int) (*model.Profile, error) {

	return nil, nil
}

func (pr *ProfileRepository) GetFullByID(id int) (*model.Profile, error) {
	nilprof := &profileSQL{}
	if err := pr.storage.pgdb.QueryRow(
		`SELECT id, name, surname, "position", description 
			FROM "shLSAIT"."objProfile"
			WHERE id = $1;`,
		id,
	).Scan(
		&nilprof.id,
		&nilprof.name,
		&nilprof.surname,
		&nilprof.position,
		&nilprof.description); err != nil {
		return nil, err
	}

	log.Println("str:59, OK")

	prof := &model.Profile{
		ID:          nilprof.id,
		Name:        nilprof.name.String,
		Surname:     nilprof.surname.String,
		Position:    nilprof.position.String,
		Description: nilprof.description.String,
		Articles:    make([]*model.Article, 0, 2),
		Drafts:      make([]*model.Article, 0, 2),
	}

	result, err := pr.storage.pgdb.Query(
		`SELECT id, title, description, date
			FROM "shLSAIT"."objArticle"
			WHERE id_auth = $1 AND article=true;`,
		id,
	)
	if err != nil {
		return nil, err
	}
	log.Println("str:79, OK")

	tmpart := &articleSQL{}

	for result.Next() {
		err := result.Scan(&tmpart.id, &tmpart.title, &tmpart.description, &tmpart.date)
		if err != nil {
			return nil, err
		}
		prof.Articles = append(prof.Articles, &model.Article{
			ID:          tmpart.id,
			Title:       tmpart.title.String,
			Description: tmpart.description.String,
			Date:        tmpart.date.String,
		})
	}
	result.Close()

	result, err = pr.storage.pgdb.Query(
		`SELECT id, title, description, date
			FROM "shLSAIT"."objArticle"
			WHERE id_auth = $1 AND article=false;`,
		id,
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&tmpart.id, &tmpart.title, &tmpart.description, &tmpart.date)
		if err != nil {
			return nil, err
		}
		prof.Drafts = append(prof.Drafts, &model.Article{
			ID:          tmpart.id,
			Title:       tmpart.title.String,
			Description: tmpart.description.String,
			Date:        tmpart.date.String,
		})
	}
	result.Close()
	return prof, nil
}
