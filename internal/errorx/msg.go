package errorx

const (
	Error                  = "未知错误!"
	SuccessMsg             = "Success!"
	ErrorParamsInvalidMsg  = "参数错误!"
	ErrorGetTaskListMsg    = "获取任务列表失败!"
	ErrorGetTaskMsg        = "获取任务失败!"
	ErrorInternalServerMsg = "内部服务器错误!"
	ErrorGetServerListMsg  = "获取边缘服务器列表失败!"
	ErrorGetClientListMsg  = "获取客户端列表失败!"
)

var MsgFlags = map[int]string{
	Err:               Error,
	Success:           SuccessMsg,
	ErrParamsInvalid:  ErrorParamsInvalidMsg,
	ErrGetTaskList:    ErrorGetTaskListMsg,
	ErrGetTask:        ErrorGetTaskMsg,
	ErrInternalServer: ErrorInternalServerMsg,
	ErrGetServerList:  ErrorGetServerListMsg,
	ErrGetClientList:  ErrorGetClientListMsg,
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Err]
}
