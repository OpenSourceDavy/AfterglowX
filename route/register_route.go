package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/repository"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewRegisterRouter(db *gorm.DB, group *gin.RouterGroup) {
	rp := repository.NewUserRepository(db)
	rc := controller.RegisterController{
		RegisterUsecase: service.NewRegisterUseCase(rp),
	}
	group.POST("/register", rc.Register)
}
