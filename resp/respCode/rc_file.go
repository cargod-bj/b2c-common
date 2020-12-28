package respCode

// 错误码，命名忽略了Error后缀
const (
	// dto映射错误
	FileDecode = prefixFile + "01"
	// 访问数据库错误
	FileDB = prefixFile + "02"
	// json解析错误
	FileJson = prefixFile + "03"
)

// 初始化 当前微服务 ResponseCode，如果要查询关联lang内容则需要在使用之前进行初始化
func InitFileRC() {
	rcMap[FileDecode] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[FileDB] = message{
		langEn: "", langId: "", langTh: "",
	}
	rcMap[FileJson] = message{
		langEn: "", langId: "", langTh: "",
	}
}
