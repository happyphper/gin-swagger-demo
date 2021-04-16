package response

const (
	Ok = 20000 // 读取成功

	AuthError        = 40100 // 账号或密码不匹配
	SmsError         = 40101 // 短信验证码错误
	CaptchaError     = 40102 // 图形验证码错误
	UsernameNotFound = 40103 // 账号不存在
	UsernameExists   = 40104 // 账号已存在
	TokenExpired     = 40110 // 身份信息失效
	ValidationError  = 42200 // 参数错误
)

type Response struct {
	// 状态码
	// 20000:成功
	// 40100:账号或密码不匹配
	// 40101:短信验证码错误
	// 40102:图形验证码错误
	// 40103:账号不存在
	// 40104:账号已存在
	// 42200:参数错误
	Code uint `json:"code" example:"20000"`
	// 消息
	Message string `json:"message" example:"ok"`
	// 数据
	Data interface{} `json:"data,omitempty"`
	// 额外数据
	Meta interface{} `json:"meta,omitempty"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    AuthError,
		Message: "ok",
		Data:    data,
	}
}

func AuthErrorRes() *Response {
	return &Response{
		Code:    AuthError,
		Message: "账号或密码不匹配",
	}
}

func TokenExpiredRes() *Response {
	return &Response{
		Code:    TokenExpired,
		Message: "身份信息失效",
	}
}

func SmsErrorRes() *Response {
	return &Response{
		Code:    SmsError,
		Message: "短信验证码错误",
	}
}

func CaptchaErrorRes() *Response {
	return &Response{
		Code:    CaptchaError,
		Message: "图形验证码错误",
	}
}

func UsernameNotFoundRes() *Response {
	return &Response{
		Code:    UsernameNotFound,
		Message: "账号不存在",
	}
}

func UsernameExistsRes() *Response {
	return &Response{
		Code:    UsernameExists,
		Message: "账号已存在",
	}
}

func ValidationRes(message string) *Response {
	return &Response{
		Code:    ValidationError,
		Message: message,
	}
}
