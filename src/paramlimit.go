package src

import "github.com/gin-gonic/gin"

// 支持context的请求key 进行限流。
// /abc?name=xxx  name 就是 key 参数
func ParamLimiter(cap int64, rate int64, key string) func(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		limiter := NewBucket(cap, rate)
		return func(c *gin.Context) {
			if c.Query(key) != "" {

				if limiter.IsAccept() {
					handler(c)
				} else {
					c.AbortWithStatusJSON(400, gin.H{"message":"fail response by param key"})
				}
			} else {
				handler(c)
			}
		}
	}
}
