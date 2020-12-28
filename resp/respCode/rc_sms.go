package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	SmsDecode = prefixSms + "01"
	// 访问数据库错误
	SmsDB = prefixSms + "02"
	// json解析错误
	SmsJson = prefixSms + "03"
)

// 初始化 当前微服务 ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitSmsRC() {
	rcMap[SmsDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[SmsDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[SmsJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
