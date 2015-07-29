package cors

import (
	"github.com/praesarium/go-engine/engine"
	"net/http"
	"strings"
)

func MiddlewareCors(origins, headers, methods []string) engine.Middleware {

	corsMethods := strings.Join(methods, ",")
	corsHeaders := strings.Join(headers, ",")

	return func(c *engine.Context) {

		corsOrigins := c.Request.Header.Get("Origin")
		if origins[0] == "*" {
			corsOrigins = "*"
		} else {

			var originFound bool
			for _, o := range origins {
				if o == corsOrigins {
					originFound = true
					break
				}
			}

			if originFound == false {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin",  corsOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Headers", corsHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", corsMethods)

		// handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.NextMiddleware()
	}
}
