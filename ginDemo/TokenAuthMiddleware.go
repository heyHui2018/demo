package demo

import (
	"os"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func responseWithError(code int, msg string, c *gin.Context) {
	resp := map[string]string{"error": msg}
	c.JSON(code, resp)
	c.Abort()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")
		if token == "" {
			responseWithError(401, "api token required", c)
			return
		}
		if token != os.Getenv("TEMP") {
			responseWithError(401, "invalid api token", c)
			return
		}
		c.String(http.StatusOK, "hello world")
		c.Writer.Header().Set("X-Request-Id", string(time.Second))
		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(TokenAuthMiddleware())
	r.Run(":12345")
}
