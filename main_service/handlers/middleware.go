package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/transactional_outbox_pattern/main_service/limiter"
)


func RateLimiterMiddleWare(context *gin.Context) {

	limiter := limiter.NewRateLimiter()

	if limiter.IsBucketEmpty() {
		context.Abort()
	}

	context.Next()
}