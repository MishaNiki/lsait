package combatstorage

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type TokenRepository struct {
	storage *CombatStorage
}

// Create ...
func (tr *TokenRepository) Create(key string, value int) error {
	err := tr.storage.redisClient.Set("tk-"+key, value, 360*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (tr *TokenRepository) Get(key string) (int, bool) {
	val, err := tr.storage.redisClient.Get("tk-" + key).Result()
	if err == redis.Nil {
		return 0, false
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return v, true
}

// Delete ...
func (tr *TokenRepository) Delete(key string) {
	tr.storage.redisClient.Del("tk-" + key)
}
