package cors

import (
	"github.com/praesarium/go-engine/engine"
	"net/http"
)

func MiddlewareCors(origin, headers, methods string) engine.Middleware {
	return func(c *engine.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
			c.Writer.Header().Set("Access-Control-Allow-Origin",  origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		}

		// handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.NextMiddleware()
	}
}
