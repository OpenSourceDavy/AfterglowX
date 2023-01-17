package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
	"github.com/wwkeyboard/sunsetwx/util"
)

type qualityUseCase struct {
}

func NewQualityUseCase() domain.QualityUseCase {
	return &qualityUseCase{}
}

func (qu *qualityUseCase) GetSunsetQuality(lat float64, lon float64) (float64, error) {
	param := make(map[string]string)
	param["lat"] = fmt.Sprintf("%f", lat)
	param["lon"] = fmt.Sprintf("%f", lon)
	param["appid"] = "b51325554e1adbdfb37ec4cbed1dcfd5"
	param["units"] = "metric"
	url := "https://pro.openweathermap.org/data/2.5/forecast/hourly"

	responseBody, err := util.QueryAPI(param, http.MethodGet, url)
	if err != nil {
		logs.Log.Error("GetMetrics OpenWeatherMap error, error message: %s", err)
	}

	var openWeatherResp domain.OpenWeatherResp
	if err := json.Unmarshal(responseBody, &openWeatherResp); err != nil {
		logs.Log.Error("GetMetrics Unmarshal error, error message: %s", err)
	}

	return 0.9, nil
}

func (qu *qualityUseCase) GetSunriseQuality(lat float64, lon float64) (float64, error) {
	return 0.9, nil
}
