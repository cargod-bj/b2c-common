package respCode

// 请求成功
const SUCCESS = "0000"

// 请求成功通用提示
const SUCCESS_MSG = "success"

// -----------------通用错误------------------------

// 未知错误
const FAILED_UNKNOWN = "9999"

// 请求失败通用提示
const FAILED_UNKNOWN_MSG = "failed unknown"

// -----------------数据库错误------------------------

// 数据库操作错误
const FAILED_DATABASE = "7001"

// 数据库操作失败通用提示
const FAILED_DATABASE_MSG = "数据库操作失败"

// -----------------请求参数错误------------------------

// 入参错误
const FAILED_PARAMS = "6001"

// 入参错误通用提示
const FAILED_PARAMS_MSG = "请求参数错误"

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

// 没有权限
const FAILED_AUTHORITY_NONE = "4001"

// 没有权限
const FAILED_AUTHORITY_NONE_MSG = "权限错误"

// 状态错误
const FAILED_STATUS_ERR = "4002"

// 状态错误
const FAILED_STATUS_ERR_MSG = "当前状态不支持此操作"
