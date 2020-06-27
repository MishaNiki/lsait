package combatstorage

import "github.com/MishaNiki/lsait/backend/internal/app/model"

type ThemeRepository struct {
	storage *CombatStorage
}

func (tr *ThemeRepository) GetByID(id int) (*model.Theme, error) {
	return nil, nil
}

func (tr *ThemeRepository) GetByUUID(uuid string) (*model.Theme, error) {
	return nil, nil
}
