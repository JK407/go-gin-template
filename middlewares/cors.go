package middlewares

import (
	"gin-template/utils/result"

	"github.com/gin-gonic/gin"
)

// Cors 跨域设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		method := c.Request.Method
		origin := c.GetHeader("Origin") // 请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 允许类型校验
		if method == "OPTIONS" {
			result.Success.ToJson(c)
			c.Abort()
			return
		}

		/*defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n", c)
				log.Printf("panic: %v\n", err)
				// log.Logger.Error("HttpError", zap.Any("HttpError", err))
			}
		}()*/

		c.Next()
		// 请求后

	}
}
