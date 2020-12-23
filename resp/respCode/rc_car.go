package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	CarDecode = prefixCar + "01"
	// 访问数据库错误
	CarDB = prefixCar + "02"
	// json解析错误
	CarJson = prefixCar + "03"
)

// 初始化 当前微服务 的ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitCarRC() {
	rcMap[CarDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[CarDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[CarJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
