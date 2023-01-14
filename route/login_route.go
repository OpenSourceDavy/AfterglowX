package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/repository"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewLoginRouter(db *gorm.DB, group *gin.RouterGroup) {
	rp := repository.NewUserRepository(db)
	lc := controller.LoginController{
		LoginUseCase: service.NewLoginUseCase(rp),
	}
	group.POST("/login", lc.Login)
}
