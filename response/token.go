package response

type TokenModel struct {
	// JWT Bearer Token
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	// 过期时间，时间戳
	ExpiresIn int64 `json:"expires_in" example:"1618535532"`
}

func (t *TokenModel) ToResponse() *Response {
	return &Response{
		Code:    Ok,
		Message: "登录成功",
		Data:    t,
		Meta:    nil,
	}
}
