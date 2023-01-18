package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/middleware"
	"github.com/wwkeyboard/sunsetwx/settings"
)

func Setup(routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	// All Public APIs
	NewRegisterRouter(publicRouterV1)
	NewLoginRouter(publicRouterV1)
	NewRefreshTokenRouter(publicRouterV1)

	protectedRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	protectedRouterV1.Use(middleware.JwtAuthMiddleware(settings.JwtSecret))
	// All Private APIs
	NewAlarmRuleRouter(protectedRouterV1)
	NewQualityRouter(protectedRouterV1)

}
