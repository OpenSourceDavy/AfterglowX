package service

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
	"github.com/wwkeyboard/sunsetwx/settings"
	"github.com/wwkeyboard/sunsetwx/util"
)

type qualityUseCase struct {
}

func NewQualityUseCase() domain.QualityUseCase {
	return &qualityUseCase{}
}

func GetMetrics(lat float64, lon float64, mType string) (domain.Metrics, error) {
	var windyResp domain.WindyResp
	var openWeatherResp domain.OpenWeatherResp

	params := []string{"temp", "dewpoint", "lclouds", "mclouds", "hclouds"}
	levels := []string{"surface"}
	windyReq := domain.WindyReq{
		Lat:        fmt.Sprintf("%f", lat),
		Lon:        fmt.Sprintf("%f", lon),
		Model:      "gfs",
		Parameters: params,
		Levels:     levels,
		Key:        settings.WindyKey,
	}

	windyReqJSON, err := json.Marshal(windyReq)
	if err != nil {
		logs.Log.Error("GetMetrics Marshal error, error message: %s", err)
	}

	responseBody, err := util.PostAPI(windyReqJSON, settings.WindyURL)
	if err != nil {
		logs.Log.Error("GetMetrics Windy error, error message: %s", err)
	}

	err = json.Unmarshal(responseBody, &windyResp)
	if err != nil {
		logs.Log.Error("GetMetrics Unmarshal error, error message: %s", err)
	}

	param := make(map[string]string)
	param["lat"] = fmt.Sprintf("%f", lat)
	param["lon"] = fmt.Sprintf("%f", lon)
	param["appid"] = settings.OpenWeatherAppID
	param["units"] = "metric"

	responseBody, err = util.GetAPI(param, settings.OpenWeatherURL)
	if err != nil {
		logs.Log.Error("GetMetrics OpenWeatherMap error, error message: %s", err)
	}

	err = json.Unmarshal(responseBody, &openWeatherResp)
	if err != nil {
		logs.Log.Error("GetMetrics Unmarshal error, error message: %s", err)
	}

	var sunSetRiseIdx_2 int
	minGap := int64(800000)
	for i, item := range windyResp.Ts {
		var gap int64
		if mType == "sunset" {
			gap = absInt(item/1000 - openWeatherResp.City.Sunset)
		}

		if mType == "sunrise" {
			gap = absInt(item/1000 - openWeatherResp.City.Sunrise)
		}

		if gap < minGap {
			sunSetRiseIdx_2 = i
			minGap = gap
		}
	}

	// returns metrics at sunrise/sunset
	res := domain.Metrics{
		Longitude:        lon,
		Latitude:         lat,
		Sunrise:          openWeatherResp.City.Sunrise,
		Sunset:           openWeatherResp.City.Sunset,
		Temperature:      windyResp.TempSurface[sunSetRiseIdx_2],
		DewPoint:         windyResp.DewpointSurface[sunSetRiseIdx_2],
		LowCloudCoverage: windyResp.LcloudsSurface[sunSetRiseIdx_2],
		MidCloudCoverage: windyResp.McloudsSurface[sunSetRiseIdx_2],
		HiCloudCoverage:  windyResp.HcloudsSurface[sunSetRiseIdx_2],
	}

	return res, err
}

func (qu *qualityUseCase) GetSunsetQuality(lat float64, lon float64) (float64, error) {

	metrics, err := GetMetrics(lat, lon, "sunset")
	if err != nil {
		logs.Log.Error("GetSunsetQuality error, error message: %s", err)
	}

	// Cloud Base Height (h) = (Temp - Dew Point ) / 10 * 1247 + Elevation
	h := (metrics.Temperature-metrics.DewPoint)/10*1247 + float64(metrics.GroundLevel)

	// Longititude offset
	lonOffset := 2 * math.Acos(6371.0088/(6371.0088+h))

	var newLat float64
	if lat < 0 && lat-lonOffset < -180 {
		newLat = 180 + lat - lonOffset
	} else {
		newLat = lat - lonOffset
	}

	metrics_2, err := GetMetrics(newLat, lon, "sunset")
	if err != nil {
		logs.Log.Error("GetSunsetQuality error, error message: %s", err)
	}

	var quality float64
	if h <= 2000 {
		quality = metrics.LowCloudCoverage - metrics_2.LowCloudCoverage
	} else if h <= 7000 {
		quality = metrics.MidCloudCoverage - metrics_2.MidCloudCoverage
	} else if h <= 12000 {
		quality = metrics.HiCloudCoverage - metrics_2.HiCloudCoverage
	}

	if quality < 0 {
		return 0, nil
	}

	return quality, nil
}

func (qu *qualityUseCase) GetSunriseQuality(lat float64, lon float64) (float64, error) {
	return 0.9, nil
}

func absInt(x int64) int64 {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}
