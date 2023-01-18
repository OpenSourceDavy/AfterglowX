package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewQualityRouter(group *gin.RouterGroup) {
	qc := controller.QualityController{
		QualityUseCase: service.NewQualityUseCase(),
	}
	group.POST("/quality", qc.GetQuality)
}
