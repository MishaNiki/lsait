package storage

import "github.com/MishaNiki/lsait/backend/internal/app/model"

type Storage interface {
	ConfirmKey() ConfirmKeyRepository
	Token() TokenRepository
	Auth() AuthRepository
	Profile() ProfileRepository
	Section() SectionRepository
	Theme() ThemeRepository
	Article() ArticleRepository
	//Draft() DraftRepository
}

// ConfirmKeyRepository ...
type ConfirmKeyRepository interface {
	Create(string, int) error
	Get(string) (int, bool)
}

// TokenRepository ...
type TokenRepository interface {
	Create(string, int) error
	Get(string) (int, bool)
	Delete(string)
}

// AuthRepository ...
type AuthRepository interface {
	Create(*model.Auth) error
	Update(*model.Auth)
	GetByID(int) (*model.Auth, error)
	GetByEmail(string) (*model.Auth, error)
	DeleteByID(int)
}

// ProfileRepository ...
type ProfileRepository interface {
	Create(*model.Profile) error
	Update(*model.Profile)
	GetByID(int) (*model.Profile, error)
	GetFullByID(int) (*model.Profile, error)
}

// SectionRepository ...
type SectionRepository interface {
	GetAll() ([]*model.Section, error)
	GetSectionsAndThemes() ([]*model.Section, error)
	GetByID(int) (*model.Section, error)
	GetByUUID(string) (*model.Section, error)
}

// ThemeRepository ...
type ThemeRepository interface {
	GetByID(int) (*model.Theme, error)
	GetByUUID(string) (*model.Theme, error)
}

// ArticleRepository ...
type ArticleRepository interface {
	Create(*model.Article)
	CreateDraft(*model.Article) error
	Update(*model.Article)
	GetEditByID(int, int) (*model.Article, error)
	GetByID(int) (*model.Article, error)
	DeleteByID(int, int)
}

// DraftRepository ...
// type DraftRepository interface {
// 	Create(*model.Article) error
// 	Update(*model.Article)
// 	GetByID(int, int) (*model.Article, error)
// 	DeleteByID(int, int)
// }
