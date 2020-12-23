package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	TranDecode = prefixCar + "01"
	// 访问数据库错误
	TranDB = prefixCar + "02"
	// json解析错误
	TranJson = prefixCar + "03"
)

// 初始化 当前微服务 ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitTransactionRC() {
	rcMap[TranDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[TranDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[TranJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
