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

func GetMetrics(lat float64, lon float64, mType string) (domain.Metrics, error) {
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

	var metric domain.Metrics
	metric.Longitude = float64(lon)
	metric.Latitude = float64(lat)
	metric.City = openWeatherResp.City.Name
	metric.Country = openWeatherResp.City.Country

	var sunSetRiseIdx int
	minGap := int64(800000)
	var totalWindSpeed float64
	var totalWindDirection int64
	for i, item := range openWeatherResp.MetricList {
		var gap int64
		if mType == "sunset" {
			gap = absInt(item.DT - openWeatherResp.City.Sunset)
		}

		if mType == "sunrise" {
			gap = absInt(item.DT - openWeatherResp.City.Sunrise)
		}

		if gap < minGap {
			sunSetRiseIdx = i
			minGap = gap
		}
		totalWindSpeed += item.Wind.Speed
		totalWindDirection += item.Wind.Deg
	}

	metric.MainDescription = openWeatherResp.MetricList[sunSetRiseIdx].Weather[0].MainDescription
	metric.SubDescription = openWeatherResp.MetricList[sunSetRiseIdx].Weather[0].SubDescription
	metric.Sunrise = openWeatherResp.City.Sunrise
	metric.Sunset = openWeatherResp.City.Sunset
	metric.Temperature = openWeatherResp.MetricList[sunSetRiseIdx].Main.Temperature
	metric.Pressure = openWeatherResp.MetricList[sunSetRiseIdx].Main.Pressure
	metric.Humidity = openWeatherResp.MetricList[sunSetRiseIdx].Main.Humidity

	// what is the default value?
	precipitation := openWeatherResp.MetricList[sunSetRiseIdx].Rain.OneHour
	if precipitation > 0 {
		metric.Precipitation = precipitation
	}
	metric.WindSpeed = openWeatherResp.MetricList[sunSetRiseIdx].Wind.Speed
	metric.AvgWindSpeed = totalWindSpeed / float64(openWeatherResp.Count)
	metric.WindDirection = openWeatherResp.MetricList[sunSetRiseIdx].Wind.Deg
	metric.AvgWindDirection = totalWindDirection / int64(openWeatherResp.Count)
	metric.CloudCoverage = openWeatherResp.MetricList[sunSetRiseIdx].Clouds.All
	metric.Visibility = openWeatherResp.MetricList[sunSetRiseIdx].Visibility

	param = make(map[string]string)
	lats := fmt.Sprintf("%f", lat)
	lons := fmt.Sprintf("%f", lon)
	param["token"] = "56ff933debba640135fa45e9d61713d8d786af63"
	url = "https://api.waqi.info/feed/geo:" + lats + ";" + lons + "/"

	responseBody, err = util.QueryAPI(param, http.MethodGet, url)
	if err != nil {
		logs.Log.Error("GetMetrics AQI error, error message: %s", err)
	}

	var aqiResp domain.AQIResp
	if err := json.Unmarshal(responseBody, &aqiResp); err != nil {
		logs.Log.Error("GetMetrics AQI Unmarshal error, error message: %s", err)
	}

	metric.AirQuality = aqiResp.Data.Aqi

	return metric, nil
}

func (qu *qualityUseCase) GetSunsetQuality(lat float64, lon float64) (float64, error) {

	// metrics, err := GetMetrics(lat, lon, "sunset")
	// if err != nil {
	// 	logs.Log.Error("GetSunsetQuality error, error message: %s", err)
	// }

	// Actual Vapor Pressure (e) = 6.11 * 10 * (7.5 * td) / (237.3 + td)

	// Standard Vapor Pressure (es) = 6.11 * 10 * (7.5 * t) / (237.3 + t)

	// Relative Humidity (rh) = e / es * 100

	// alpha(t, rh) = ln(rh / 100) + (a * t) / (b + t)

	// Dew Point (td) = (b * alpha(t, rh)) / (a - alpha(t, rh))

	// Cloud Base Height (h) = (Temp - Dew Point ) / 10 * 1247 + Elevation

	return 0.9, nil
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
