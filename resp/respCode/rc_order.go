package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	OrderDecode = prefixOrder + "01"
	// 访问数据库错误
	OrderDB = prefixOrder + "02"
	// json解析错误
	OrderJson = prefixOrder + "03"
)

// 初始化 当前微服务 的ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitOrderRC() {
	rcMap[OrderDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[OrderDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[OrderJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
