package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/repository"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewRefreshTokenRouter(group *gin.RouterGroup) {
	ur := repository.NewUserRepository()
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: service.NewRefreshTokenUsecase(ur),
	}
	group.POST("/refreshToken", rtc.RefreshToken)
}
