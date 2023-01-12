package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/settings"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	gin.SetMode(settings.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/quality", GetSunsetQuality)

		apiv1.POST("/register", CreateUser)

		apiv1.POST("/login", UserLogin)

		apiv1.POST("/rules", GetAlarmRules)

		apiv1.POST("/create_rule", CreateAlarmRule)

		apiv1.POST("/update_rule", UpdateAlarmRule)

		apiv1.POST("/delete_rule", DeleteAlarmRule)
	}

	return r
}
