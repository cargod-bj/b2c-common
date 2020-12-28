package respCode

import i18nKits "github.com/cargod-bj/b2c-common/i18n"

// 各个微服务的前缀
const (
	prefixCommon      = "00"
	prefixGateway     = "01"
	prefixCar         = "02"
	prefixOrder       = "03"
	prefixTransaction = "04"
	prefixCustomer    = "05"
	prefixSms         = "06"
	prefixStaff       = "07"
	prefixStore       = "08"
	prefixFile        = "09"
)

const (
	// 通用成功
	RcSuccess = prefixCommon + "00"
)

var rcMap = make(map[string]message)

// 消息体，对应不同语言
type message struct {
	langEn string
	langId string
	langTh string
	langZh string
}

// 初始化所有微服务的ResponseCode，如果要查询关联lang内容则需要先调用本类方法进行初始化。
// 如果是在微服务中使用，只需要调用相关微服务的Init方法即可，本方法是全量初始化方法。
func InitRC() {
	rcMap[RcSuccess] = message{
		langEn: "Success", langId: "Success Id", langTh: "Success Th", langZh: "Success Zh",
	}

	InitCarRC()
	InitCustomerRC()
	InitFileRC()
	InitGatewayRC()
	InitSmsRC()
	InitStaffRC()
	InitStoreRC()
	InitTransactionRC()
}

// 根据responseCode获取responseMessage，使用当前默认语言
// 使用本方法前必须先调用相应微服务的Init类方法
func GetRM(code string) string {
	return GetRMByLang(code, i18nKits.GetLang())
}

// 根据responseCode获取responseMessage，使用lang指定语言
// 使用本方法前必须先调用相应微服务的Init类方法
func GetRMByLang(code, lang string) string {
	if code == "" {
		return "Failed unknown."
	}
	item, ok := rcMap[code]
	if !ok {
		return "Failed unknown"
	}

	switch lang {
	case i18nKits.LangEn:
		return item.langEn
	case i18nKits.LangId:
		return item.langId
	case i18nKits.LangTh:
		return item.langTh
	case i18nKits.LangZh:
		return item.langZh
	}

	return item.langEn
}
