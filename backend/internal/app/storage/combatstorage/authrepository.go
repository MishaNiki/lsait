package combatstorage

import "github.com/MishaNiki/lsait/backend/internal/app/model"

type AuthRepository struct {
	storage *CombatStorage
}

func (ar *AuthRepository) Create(a *model.Auth) error {
	if err := ar.storage.pgdb.QueryRow(
		"INSERT INTO \"shLSAIT\".\"objAuth\" (email, password) Values($1, $2) RETURNING id",
		a.Email,
		a.EncryptedPassword,
	).Scan(&a.ID); err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepository) Update(a *model.Auth) {
	ar.storage.pgdb.QueryRow(
		"UPDATE \"shLSAIT\".\"objAuth\" SET email=$1, password=$2 WHERE id = $3",
		a.Email,
		a.EncryptedPassword,
		a.ID,
	)
}

func (ar *AuthRepository) GetByID(id int) (*model.Auth, error) {
	a := &model.Auth{}
	if err := ar.storage.pgdb.QueryRow(
		"SELECT id, email, password FROM \"shLSAIT\".\"objAuth\" WHERE id = $1",
		id,
	).Scan(
		&a.ID,
		&a.Email,
		&a.EncryptedPassword); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *AuthRepository) GetByEmail(email string) (*model.Auth, error) {
	a := &model.Auth{}
	if err := ar.storage.pgdb.QueryRow(
		"SELECT id, email, password FROM \"shLSAIT\".\"objAuth\" WHERE email = $1",
		email,
	).Scan(
		&a.ID,
		&a.Email,
		&a.EncryptedPassword); err != nil {
		return nil, err
	}
	return a, nil
}

// DeleteByID ...
func (ar *AuthRepository) DeleteByID(id int) {
	ar.storage.pgdb.QueryRow(
		"DELETE FROM \"shLSAIT\".\"objAuth\" WHERE id = $1",
		id,
	)
}
