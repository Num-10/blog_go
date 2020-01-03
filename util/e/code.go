package e

const(
	SERVICE_SUCCESS = 200
	SERVICE_FIAL = 400
	PRRAMS_ERROR = 401
	SERVICE_ERROR = 500
	SERVICE_CONNECT_MODEL = 501
	READ_CONFIG_ERROR = 510

	TOKEN_IN_VAIN = 1001
	TOKEN_CREATE_FAIL = 1002
	IMAGE_OVER_SIZE = 1003
	IMAGE_NOT_ALLOW_EXT = 1004
	IMAGE_SAVE_FIAL = 1005

	LOGIN_PARAM_EMPTY = 2001
	LOGIN_PARAM_ERROR = 2002

	TITLE_IS_EXISTS = 2003
)

var Message = map[int]string {
	SERVICE_SUCCESS         : "操作成功",
	SERVICE_FIAL			: "操作失败，请稍后重试！",
	PRRAMS_ERROR			: "参数有误",
	SERVICE_ERROR           : "error",
	SERVICE_CONNECT_MODEL   : "connect model fail",
	READ_CONFIG_ERROR       : "load config fail",

	TOKEN_IN_VAIN           : "token in vain",
	TOKEN_CREATE_FAIL       : "create token fail",

	IMAGE_OVER_SIZE			: "图片超过大小限制",
	IMAGE_NOT_ALLOW_EXT		: "图片格式不支持",
	IMAGE_SAVE_FIAL			: "图片上传失败",

	LOGIN_PARAM_EMPTY       : "账号或密码不能为空",
	LOGIN_PARAM_ERROR       : "账号或密码错误",

	TITLE_IS_EXISTS			: "标签名称已存在",
}

func GetMsg(code int) (string) {
	msg, ok := Message[code]
	if ok {
		return msg
	}
	return Message[SERVICE_ERROR]
}

