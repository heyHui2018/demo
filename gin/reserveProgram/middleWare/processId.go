package middleWare

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/demo/gin/reserveProgram/utils"
	"time"
)

func GenerateProcessId() gin.HandlerFunc {
	return func(c *gin.Context) {
		processId := time.Now().Format("20060102150405") + utils.GetRandString()
		c.Set("processId", processId)
		c.Next()
	}
}
