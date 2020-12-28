package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	StoreDecode = prefixStore + "01"
	// 访问数据库错误
	StoreDB = prefixStore + "02"
	// json解析错误
	StoreJson = prefixStore + "03"
)

// 初始化 当前微服务 ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitStoreRC() {
	rcMap[StoreDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[StoreDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[StoreJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
