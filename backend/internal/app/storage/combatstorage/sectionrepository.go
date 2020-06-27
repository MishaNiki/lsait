package combatstorage

import (
	"database/sql"

	"github.com/MishaNiki/lsait/backend/internal/app/model"
)

type SectionRepository struct {
	storage *CombatStorage
}

func (sr *SectionRepository) GetAll() ([]*model.Section, error) {
	result := make([]*model.Section, 0, 2)
	rows, err := sr.storage.pgdb.Query("SELECT * FROM \"shLSAIT\".\"objSection\"")
	if err != nil {
		return nil, err
	}
	var (
		id    int
		uuid  string
		title string
	)
	for rows.Next() {
		err := rows.Scan(&id, &uuid, &title)
		if err != nil {
			return nil, err
		}
		result = append(result, &model.Section{ID: id, UUID: uuid, Title: title})
	}
	rows.Close()
	return result, nil
}

// GetSectionsAndThemes ...
func (sr *SectionRepository) GetSectionsAndThemes() ([]*model.Section, error) {
	sections := make([]*model.Section, 0, 2)

	result, err := sr.storage.pgdb.Query(
		`SELECT sec.id, sec.title, th.id, th.title  
			FROM "shLSAIT"."objSection" sec
			FULL JOIN "shLSAIT"."objTheme" th
			ON sec.id = th.id_section
			ORDER BY sec.id, th.position;`,
	)
	if err != nil {
		return nil, err
	}

	var i, curSec int = -1, -1
	row := struct {
		idSec    int
		titleSec string
		idTh     sql.NullInt32
		titleTh  sql.NullString
	}{}

	for result.Next() {
		err := result.Scan(&row.idSec, &row.titleSec, &row.idTh, &row.titleTh)
		if err != nil {
			return nil, err
		}
		if curSec == row.idSec {
			if row.idTh.Valid {
				sections[i].Themes = append(sections[i].Themes, &model.Theme{
					ID:    int(row.idTh.Int32),
					Title: row.titleTh.String,
				})
			}
		} else {
			curSec = row.idSec
			i++
			sections = append(sections, &model.Section{
				ID:     row.idSec,
				Title:  row.titleSec,
				Themes: make([]*model.Theme, 0, 2),
			})
			if row.idTh.Valid {
				sections[i].Themes = append(sections[i].Themes, &model.Theme{
					ID:    int(row.idTh.Int32),
					Title: row.titleTh.String,
				})
			}
		}
	}
	return sections, nil
}

func (sr *SectionRepository) GetByID(int) (*model.Section, error) {

	return nil, nil
}

// GetByUUID ...
func (sr *SectionRepository) GetByUUID(uuid string) (*model.Section, error) {
	sec := &model.Section{UUID: uuid}
	if err := sr.storage.pgdb.QueryRow(
		"SELECT id, title FROM \"shLSAIT\".\"objSection\" WHERE uuid = $1",
		uuid,
	).Scan(
		&sec.ID,
		&sec.Title); err != nil {
		return nil, err
	}
	sec.Themes = make([]*model.Theme, 0, 2)
	result, err := sr.storage.pgdb.Query(
		`SELECT th.id, th.title, art.id, art.title, art.description
			FROM "shLSAIT"."objTheme" th
			FULL JOIN "shLSAIT"."viewArticle" art 
			ON art.id_theme = th.id
			WHERE th.id_section = $1
			ORDER BY th.position;`,
		sec.ID,
	)
	if err != nil {
		return nil, err
	}
	var (
		i              int = -1
		curentID       int
		idTheme        int
		titleTheme     string
		idArt          sql.NullInt32
		titleArt       sql.NullString
		descriptionArt sql.NullString
	)
	for result.Next() {
		err := result.Scan(&idTheme, &titleTheme, &idArt, &titleArt, &descriptionArt)
		if err != nil {
			return nil, err
		}

		if curentID == idTheme {
			if idArt.Valid {
				sec.Themes[i].Articles = append(sec.Themes[i].Articles, &model.Article{
					ID:          int(idArt.Int32),
					Title:       titleArt.String,
					Description: descriptionArt.String,
				})
			}
		} else {
			sec.Themes = append(sec.Themes, &model.Theme{
				ID:       idTheme,
				Title:    titleTheme,
				Articles: make([]*model.Article, 0, 2),
			})
			i++
			if idArt.Valid {
				sec.Themes[i].Articles = append(sec.Themes[i].Articles, &model.Article{
					ID:          int(idArt.Int32),
					Title:       titleArt.String,
					Description: descriptionArt.String,
				})
			}
			curentID = idTheme
		}
	}
	result.Close()
	return sec, nil
}
