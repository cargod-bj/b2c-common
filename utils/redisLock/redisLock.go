package redisLock

import (
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client

/**
 *	Redis锁 result true 加锁成功，false已存在锁
 */
func Lock(lockKey string, expiration time.Duration) (result bool) {
	if expiration == 0 {
		expiration = time.Second * time.Duration(30)
	}
	var lock = redisClient.SetNX(lockKey, lockKey, expiration)
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
func InitLockUtil(client *redis.Client) {
	redisClient = client
}
