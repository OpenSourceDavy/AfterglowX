package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/repository"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewRegisterRouter(group *gin.RouterGroup) {
	rp := repository.NewUserRepository()
	rc := controller.RegisterController{
		RegisterUsecase: service.NewRegisterUseCase(rp),
	}
	group.POST("/register", rc.Register)
}
