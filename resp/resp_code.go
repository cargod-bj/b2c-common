package resp

// 请求成功
const SUCCESS = "0000"

// 请求成功通用提示
const SUCCESS_MSG = "success"

// -----------------通用错误------------------------

// 未知错误
const FAILED_UNKNOWN = "9999"

// 请求失败通用提示
const FAILED_UNKNOWN_MSG = "failed unknown"

// -----------------代码报错，比如空指针之类的------------------------

// dto映射失败
const FAILED_DTO_DECODE = "5001"

// dto映射失败提示
const FAILED_DTO_DECODE_MSG = "数据映射失败"

// dto映射失败：data=nil
const FAILED_DTO_DATA_NIL = "5002"

// dto映射失败提示：data=nil
const FAILED_DTO_DATA_NIL_MSG = "Data为空"

// -----------------常用业务错误------------------------
