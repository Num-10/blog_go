package e

const(
	SERVICE_SUCCESS = 200
	SERVICE_ERROR = 500
	SERVICE_CONNECT_MODEL = 501
	READ_CONFIG_ERROR = 510

	TOKEN_IN_VAIN = 1001
	TOKEN_CREATE_FAIL = 1002

	LOGIN_PARAM_EMPTY = 2001
	LOGIN_PARAM_ERROR = 2002
)

var Message = map[int]string {
	SERVICE_SUCCESS         : "ok",
	SERVICE_ERROR           : "fail",
	SERVICE_CONNECT_MODEL   : "connect model fail",
	READ_CONFIG_ERROR       : "load config fail",

	TOKEN_IN_VAIN           : "token in vain",
	TOKEN_CREATE_FAIL       : "create token fail",

	LOGIN_PARAM_EMPTY       : "账号或密码不能为空",
	LOGIN_PARAM_ERROR       : "账号或密码错误",
}

func GetMsg(code int) (string) {
	msg, ok := Message[code]
	if ok {
		return msg
	}
	return Message[SERVICE_ERROR]
}

