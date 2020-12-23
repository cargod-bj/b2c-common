package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	StaffDecode = prefixCar + "01"
	// 访问数据库错误
	StaffDB = prefixCar + "02"
	// json解析错误
	StaffJson = prefixCar + "03"
)

// 初始化 当前微服务 ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitStaffRC() {
	rcMap[StaffDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[StaffDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[StaffJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
