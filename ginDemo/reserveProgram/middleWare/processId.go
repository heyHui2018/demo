package middleWare

import (
	"github.com/heyHui2018/demo/ginDemo/reserveProgram/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func GenerateProcessId() gin.HandlerFunc {
	return func(c *gin.Context) {
		processId := time.Now().Format("20060102150405") + utils.GetRandString()
		c.Set("processId", processId)
		c.Next()
	}
}
