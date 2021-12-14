# Middleware to Handle Request Timeout in Gin

## Installation

Installation

```sh
go get github.com/s-wijaya/gin-timeout
```

Import it in your code:

```go
import (
    // other imports
    timeout "github.com/s-wijaya/gin-timeout"
)
```

## Usage

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	timeout "github.com/s-wijaya/gin-timeout"
)

func  main() {
	engine := gin.New()
 
	responseBodyTimeout := gin.H{
		"code": http.StatusRequestTimeout,
		"message": "request timeout, response is sent from middleware"}

	engine.Use(timeout.TimeoutHandler(5 * time.Second, http.StatusRequestTimeout, responseBodyTimeout)

	engine.GET("/excedd", ExceedTimeout)

	engine.GET("/inrange", InRangeTimeout)
	
	engine.Run()
}

func ExceedTimeout(c *gin.Context) {
	timeout.APIWrapper(c, func(c *gin.Context) (int, interface{}) {
		time.Sleep(8 * time.Second) // all process is here, including log, calculation, retrieving data etc
		return http.StatusRequestTimeout, gin.H{"code": http.StatusRequestTimeout, "message": "exceed timeout limit"}
	})
}

func InRangeTimeout(c *gin.Context) {
	timeout.APIWrapper(c, func(c *gin.Context) (int, interface{}) {
		time.Sleep(3 * time.Second) // all process is here, including log, calculation, retrieving data etc
		return http.StatusOK, gin.H{"code": http.StatusOK, "message": "in range timeout limit"}
	})
}
```