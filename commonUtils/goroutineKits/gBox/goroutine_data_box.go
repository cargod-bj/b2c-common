package gBox

import (
	"fmt"
	"github.com/cargod-bj/b2c-common/commonUtils/goroutineKits"
	"strings"
	"sync"
	"time"
)

var gBoxCache sync.Map

// 默认的每个对象的存活时长
const defaultLifeDuration = int64(5 * time.Minute)

const gKeyHeader = "$g%d$k"

// 自定义的对象存活时长，如果 <=0 则使用defaultLifeDuration
var lifeDuration = int64(0)

var initHolder = sync.Once{}
var ticker *time.Ticker
var lock sync.Mutex

func initGBox() {
	println("尝试初始化新的回收器")
	initHolder.Do(func() {
		println("初始化新的回收器")
		lock.Lock()
		ticker = time.NewTicker(3 * time.Minute)
		lock.Unlock()
		// 定时删除缓存中无用的cache信息，缓存的声明周期是5分钟，这里的轮训是3分钟删一次，所以一个cache最大生命周期是6分钟
		go func() {
			count := 0
			for {
				select {
				case <-ticker.C:

					println("recycle goroutine box cache")

					count++

					curr := time.Now().UnixNano()
					gBoxCache.Range(func(k, v interface{}) bool {
						count = 0
						if v == nil {
							return true
						}
						l := v.(value)
						if l.recycleTime > 0 && curr > l.recycleTime {
							gBoxCache.Delete(k)
						}
						return true
					})
				}

				// 如果12分钟都没有数据，则停止自动回收
				if count > 3 {
					RecycleAll()
					break
				}
			}
		}()
	})
}

// 添加数据到当前goroutine中
// k 数据的key，可以使用Get方法获取出来。
// v 保存到当前goroutine的数据
// autoRecycle 是否自动回收，0自动回收，否则不自动回收
func Put(k string, v interface{}, autoRecycle ...uint32) {
	var duration int64
	if len(autoRecycle) > 0 && autoRecycle[0] != 0 {
		duration = 0
	} else {
		duration = getLifeDuration()
	}
	PutWithDuration(k, v, duration)
}

// 添加数据到当前goroutine中
// k 数据的key，可以使用Get方法获取出来。
// v 保存到当前goroutine的数据
// lifeDuration 对象v的生命周期，如果大于0则会按照指定时长自动回收，如果小于等于0，则不会自动回收，需要你手动回收
func PutWithDuration(k string, v interface{}, lifeDuration int64) {
	var recycleTime int64
	if lifeDuration > 0 {
		recycleTime = time.Now().UnixNano() + lifeDuration
		initGBox()
	}
	d := value{value: v, recycleTime: recycleTime}

	key := getKey(k)

	gBoxCache.Store(key, d)
}

// 获取当前goroutine下指定key的数据，参见Put方法
func Get(k string) interface{} {
	gid := goroutineKits.GetGID()
	return GetWithGid(gid, getKeyWithGid(gid, k))
}

// 获取指定goroutineId的数据
func GetWithGid(gid uint64, k string) interface{} {
	load, _ := gBoxCache.Load(getKeyWithGid(gid, k))
	return load
}

// 回收当前goroutine下指定key的数据
func Recycle(k string) {
	key := getKey(k)
	gBoxCache.Delete(key)
}

// 回收当前goroutine下的数据
func RecycleCurrent() {
	RecycleGid(goroutineKits.GetGID())
}

// 回收指定goroutine下的数据
func RecycleGid(gid uint64) {
	gHeader := fmt.Sprintf(gKeyHeader, gid)
	gBoxCache.Range(func(k, v interface{}) bool {
		if k == nil {
			return true
		}
		key := k.(string)
		if strings.Index(key, gHeader) == 0 {
			gBoxCache.Delete(k)
		}
		return true
	})
}

// 回收gBox持有的所有数据
func RecycleAll() {

	println("回收所有内容")

	gBoxCache.Range(func(k, v interface{}) bool {
		gBoxCache.Delete(k)
		return true
	})

	lock.Lock()
	defer lock.Unlock()
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}

	initHolder = sync.Once{}
}

func getKeyWithGid(gid uint64, k string) string {
	return fmt.Sprintf(gKeyHeader+"%s", gid, k)
}

func getKey(k string) string {
	gid := goroutineKits.GetGID()
	return getKeyWithGid(gid, k)
}

func getLifeDuration() int64 {
	if lifeDuration > 0 {
		return lifeDuration
	}
	return defaultLifeDuration
}

type value struct {
	value       interface{}
	recycleTime int64
}
