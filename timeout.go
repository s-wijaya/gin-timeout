package gintimeout

import "github.com/gin-gonic/gin"

type ServiceProcess func(*gin.Context) (int, interface{})

func APIWrapper(c *gin.Context, process ServiceProcess) {
	var apiResponseCode int
	var apiOutputData interface{}

	ctx := c.Request.Context()

	doneChan := make(chan bool)

	go func() {
		apiResponseCode, apiOutputData = process(c)
		close(doneChan)
	}()

	select {
	case <-ctx.Done():
		return
	case <-doneChan:
		c.JSON(apiResponseCode, apiOutputData)
	}
}
