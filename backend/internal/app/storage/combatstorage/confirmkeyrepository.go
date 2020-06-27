package combatstorage

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type ConfirmKeyRepository struct {
	storage *CombatStorage
}

// Create ...
func (ckr *ConfirmKeyRepository) Create(key string, value int) error {
	err := ckr.storage.redisClient.Set("ck-"+key, value, 360*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (ckr *ConfirmKeyRepository) Get(key string) (int, bool) {
	val, err := ckr.storage.redisClient.Get("ck-" + key).Result()
	if err == redis.Nil {
		return 0, false
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	ckr.storage.redisClient.Del("ck-" + key)
	return v, true
}
