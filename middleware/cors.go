package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"Origin"},
			//AllowCredentials: true,
			AllowAllOrigins: true,
			ExposeHeaders:   []string{"Content-Length", "Authorization"},
			MaxAge:          12 * time.Hour,
		})
	}
}
