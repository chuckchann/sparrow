package serror

var (
	SUCCESS         = &SError{Code: 0, Msg: "success"}
	ERR_SYSTEM      = &SError{Code: -1, Msg: "系统繁忙"}
	ERR_MARSHAL_ANY = &SError{Code: 999, Msg: "格式错误"}
	ERR_RATE        = &SError{Code: 1000, Msg: "访问频率过快"}
	ERR_PARAM       = &SError{Code: 1001, Msg: "参数错误"}
	ERR_AUTH        = &SError{Code: 1002, Msg: "TOKEN校验错误"}
	ERR_LOG         = &SError{Code: 1003, Msg: "日志记录请求错误"}
	ERR_FUSING      = &SError{Code: 1004, Msg: "服务熔断"}
	ERR_PARASE      = &SError{Code: 1005, Msg: "解析错误"}
	ERR_UNFOUND_CLI = &SError{Code: 1006, Msg: "无法找到客户端"}
	ERR_CONNECT     = &SError{Code: 1007, Msg: "GRPC连接失败"}
	ERR_CALL        = &SError{Code: 1007, Msg: "调用失败"}
	ERR_BIZ         = &SError{Code: 1008, Msg: "业务调用错误"}
)
