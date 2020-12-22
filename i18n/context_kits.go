package i18nKits

import (
	"fmt"
	"github.com/cargod-bj/b2c-common/commonUtils/goroutineKits"
	"strings"
	"sync"
	"time"
)

var languageCache sync.Map

const I18nContextKey = "language$Value$Key"
const defaultLifeDuration = int64(5 * time.Minute)

func init() {
	// 定时删除缓存中无用的cache信息，缓存的声明周期是5分钟，这里的轮训是3分钟删一次，所以一个cache最大生命周期是6分钟左右
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for {
			println("start recycle i18n cache")
			select {
			case <-ticker.C:
				curr := time.Now().UnixNano()
				languageCache.Range(func(key, value interface{}) bool {
					if value == nil {
						return true
					}
					l := value.(lang)
					if l.recycleTime > 0 && curr > l.recycleTime {
						languageCache.Delete(key)
					}
					return true
				})
			}
		}
	}()
}

// 初始化language，如果autoRecycle=true，
func InitI18n(acceptLanguageHeader string, autoRecycle bool) {
	switch {
	case strings.Contains(acceptLanguageHeader, LangEn):
		initInner(LangEn, autoRecycle)
	case strings.Contains(acceptLanguageHeader, LangId):
		initInner(LangId, autoRecycle)
	case strings.Contains(acceptLanguageHeader, LangTh):
		initInner(LangTh, autoRecycle)
	case strings.Contains(acceptLanguageHeader, LangZh):
		initInner(LangZh, autoRecycle)
	default:
		initInner("", autoRecycle)
	}
}

// 初始化language
func InitI18nByLang(language string, autoRecycle bool) {
	initInner(language, autoRecycle)
}

func initInner(language string, autoRecycle bool) {
	var rt int64
	if autoRecycle {
		rt = time.Now().UnixNano() + defaultLifeDuration
	}
	l := lang{lang: language, recycleTime: rt}
	if language == "" {
		l.lang = getDefaultLangByLocal()
	} else if !allSupportLang.Contain(language) {
		panic(fmt.Sprintf("暂时不支持的语言类型:%s", language))
	}
	gid := goroutineKits.GetGID()
	languageCache.Store(gid, l)
}

// 获取当前的语言
func GetLang() string {
	gid := goroutineKits.GetGID()
	load, ok := languageCache.Load(gid)
	if !ok {
		return getDefaultLangByLocal()
	}
	l := load.(lang)
	return l.lang
}

// 手动回收当前语言缓存
func Recycle() {
	gid := goroutineKits.GetGID()
	languageCache.Delete(gid)
}

type lang struct {
	lang        string
	recycleTime int64
}
