package config

const (
	Success=0 //正常返回码
	ErrUserName     = 1000 // 用户名错误
	ErrUserPassword = 1001 // 用户密码错误
	ErrSetRedisUserInfo=1002//缓存用户信息失败
	ErrAuth=1003//授权失败

	ErrSourceTimeLimit=2000//资源过期
	ErrSourceNotFound=2001//资源不存在
)
var ErrorMessage =map[int]string{
	Success : "请求正常。(200)",
	ErrUserName: "用户名错误。(1000)",
	ErrUserPassword: "用户密码错误。(1001)",
	ErrSetRedisUserInfo:"缓存用户信息失败。(1002)",
	ErrAuth:"授权失败。(1003)",
	ErrSourceTimeLimit: "资源过期。(2000)",
	ErrSourceNotFound: "资源不存在。(2001)",
}

func GetErrorMessage(code int) string {
	return ErrorMessage[code]
}