package redisLock

import (
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client

/**
 *	Redis锁 result true 加锁成功，false已存在锁
 */
func Lock(lockKey string, expiration uint64) (result bool) {
	var expTime time.Duration
	if expiration == 0 {
		expTime = time.Second * time.Duration(30)
	} else {
		expTime = time.Second * time.Duration(expiration)
	}
	var lock = redisClient.SetNX(lockKey, lockKey, expTime)
	return lock.Val()
}

//解锁  true 解锁成功 false解锁失败
func UnLock(lockKey string) (result bool) {
	var res = redisClient.Del(lockKey)
	if res.Val() == 1 {
		result = true
	} else {
		result = false
	}
	return result
}

//使用前需要先进行初始化
func InitRedisLock(client *redis.Client) {
	redisClient = client
}
