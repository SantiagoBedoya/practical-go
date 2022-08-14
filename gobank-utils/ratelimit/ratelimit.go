package ratelimit

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// GetRateLimit return Gin Middleware with memory store by default
func GetRateLimit(formatted string, store limiter.Store) gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		log.Fatal(err.Error())
	}
	if store == nil {
		store = memory.NewStore()
	}
	return mgin.NewMiddleware(limiter.New(store, rate))
}
