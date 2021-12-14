package gintimeout

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutHandler(timeout time.Duration, responseCodeTimeout int, responseBodyTimeout interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.JSON(responseCodeTimeout, responseBodyTimeout)
				c.Abort()
			}

			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
