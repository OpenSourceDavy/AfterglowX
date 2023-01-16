package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/repository"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewLoginRouter(group *gin.RouterGroup) {
	rp := repository.NewUserRepository()
	lc := controller.LoginController{
		LoginUseCase: service.NewLoginUseCase(rp),
	}
	group.POST("/login", lc.Login)
}
