package result

import "net/http"

var (
	Success = Resp(http.StatusOK, "success")                // 通用成功
	Err     = Resp(http.StatusInternalServerError, "error") // 通用错误

	ErrNotFound  = Resp(331, "未找到")
	ErrParam     = Resp(422, "参数有误")
	ErrSignParam = Resp(401, "签名参数有误")
	ErrService   = Resp(500, "服务异常")

	ErrTokenParam          = Resp(10001, "Token不能为空")
	ErrTokenInvalid        = Resp(10002, "Invalid token")
	ErrTokenExpired        = Resp(10003, "Expired token")
	ErrTokenExpiredSignOut = Resp(10004, "Expired token")
)
