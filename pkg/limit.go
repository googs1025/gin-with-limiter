package pkg

import (
	"gin-with-limiting/pkg/bucket"
	"github.com/gin-gonic/gin"
)

func Limiter(cap int64) func(handler gin.HandlerFunc) gin.HandlerFunc {

	limiter := bucket.NewBucket(10, 1)
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
