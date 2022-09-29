package src

import "github.com/gin-gonic/gin"

func Limiter(cap int) func(handler gin.HandlerFunc) gin.HandlerFunc {

	limiter := NewBucket(10)
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			if limiter.IsAccept() {
				handler(c)
			} else {
				c.AbortWithStatusJSON(400, gin.H{"message": "fail by limiter"})
			}
		}
	}

}
