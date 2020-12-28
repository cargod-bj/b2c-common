package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	CustDecode = prefixCustomer + "01"
	// 访问数据库错误
	CustDB = prefixCustomer + "02"
	// json解析错误
	CustJson = prefixCustomer + "03"
)

// 初始化 当前微服务 ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitCustomerRC() {
	rcMap[CustDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[CustDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[CustJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
