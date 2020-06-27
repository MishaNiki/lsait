package combatstorage

import (
	"database/sql"
	"fmt"

	"github.com/MishaNiki/lsait/backend/internal/app/storage"
	"github.com/go-redis/redis"

	_ "github.com/lib/pq" // driver postgresql
)

type CombatStorage struct {
	redisClient *redis.Client
	pgdb        *sql.DB

	confirmKeyRepository storage.ConfirmKeyRepository
	tokenRepository      storage.TokenRepository
	authRepository       storage.AuthRepository
	profileRepository    storage.ProfileRepository
	sectionRepository    storage.SectionRepository
	themeRepository      storage.ThemeRepository
	articleRepository    storage.ArticleRepository
}

type Config struct {
	PgUser        string `json:"pg-user"`
	PgPassword    string `json:"pg-password"`
	PgDBName      string `json:"pg-dbname"`
	PgSSLMode     string `json:"pg-sslmode"`
	RedisAddr     string `json:"r-addr"`
	RedisPassword string `json:"r-password"`
	RedisDB       int    `json:"r-db"`
}

// New ...
func New() *CombatStorage {
	return &CombatStorage{}
}

// Configure ...
func (cs *CombatStorage) Configure(config *Config) error {
	if err := cs.OpenRedit(config); err != nil {
		return err
	}
	if err := cs.OpenPostgres(config); err != nil {
		return err
	}
	return nil
}

// ConfigureRedit ...
func (cs *CombatStorage) OpenRedit(config *Config) error {

	cs.redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	if _, err := cs.redisClient.Ping().Result(); err != nil {
		return err
	}

	return nil
}

// ConfigurePostgres ...
func (cs *CombatStorage) OpenPostgres(config *Config) error {

	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		config.PgUser,
		config.PgPassword,
		config.PgDBName,
		config.PgSSLMode,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	cs.pgdb = db
	return nil
}

// Close ...
func (cs *CombatStorage) Close() {
	cs.pgdb.Close()
	cs.redisClient.Close()
}

// ConfirmKey ...
func (cs *CombatStorage) ConfirmKey() storage.ConfirmKeyRepository {
	if cs.confirmKeyRepository != nil {
		return cs.confirmKeyRepository
	}
	cs.confirmKeyRepository = &ConfirmKeyRepository{
		storage: cs,
	}
	return cs.confirmKeyRepository
}

// Token ...
func (cs *CombatStorage) Token() storage.TokenRepository {
	if cs.tokenRepository != nil {
		return cs.tokenRepository
	}
	cs.tokenRepository = &TokenRepository{
		storage: cs,
	}
	return cs.tokenRepository
}

// Auth ...
func (cs *CombatStorage) Auth() storage.AuthRepository {
	if cs.authRepository != nil {
		return cs.authRepository
	}
	cs.authRepository = &AuthRepository{
		storage: cs,
	}
	return cs.authRepository
}

// Profile ...
func (cs *CombatStorage) Profile() storage.ProfileRepository {
	if cs.profileRepository != nil {
		return cs.profileRepository
	}
	cs.profileRepository = &ProfileRepository{
		storage: cs,
	}
	return cs.profileRepository
}

// Section ...
func (cs *CombatStorage) Section() storage.SectionRepository {
	if cs.sectionRepository != nil {
		return cs.sectionRepository
	}
	cs.sectionRepository = &SectionRepository{
		storage: cs,
	}
	return cs.sectionRepository
}

// Theme ...
func (cs *CombatStorage) Theme() storage.ThemeRepository {
	if cs.themeRepository != nil {
		return cs.themeRepository
	}
	cs.themeRepository = &ThemeRepository{
		storage: cs,
	}
	return cs.themeRepository
}

// Article ...
func (cs *CombatStorage) Article() storage.ArticleRepository {
	if cs.articleRepository != nil {
		return cs.articleRepository
	}
	cs.articleRepository = &ArticleRepository{
		storage: cs,
	}
	return cs.articleRepository
}

// Draft ...
// func (cs *CombatStorage) Draft() storage.DraftRepository {
// 	if cs.draftRepository != nil {
// 		return cs.draftRepository
// 	}
// 	cs.draftRepository = &DraftRepository{
// 		storage: cs,
// 	}
// 	return cs.draftRepository
// }
