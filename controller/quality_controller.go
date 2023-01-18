package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
)

type QualityController struct {
	QualityUseCase domain.QualityUseCase
}

func (qc *QualityController) GetQuality(c *gin.Context) {
	var req domain.QualityRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var quality float64

	if req.Type == "sunset" {
		quality, err = qc.QualityUseCase.GetSunsetQuality(req.Lat, req.Lon)
	}

	if req.Type == "sunrise" {
		quality, err = qc.QualityUseCase.GetSunriseQuality(req.Lat, req.Lon)
	}

	if err != nil {
		logs.Log.Error("GetQuality controller error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.QualityResponse{
		Code:    1,
		Message: "success",
		Data:    quality,
	})
}
