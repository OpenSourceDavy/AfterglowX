package quality

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// sunset time, sunrise time,
// cloud coverage, cloud height
// precipitation, humidity,
// wind speed, wind direction,
// air quality

// https://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={API key}

type Metrics struct {
	Longitude        float64
	Latitude         float64
	City             string
	Country          string
	MainDescription  string
	SubDescription   string
	Sunrise          int64
	Sunset           int64
	Temperature      float64
	Pressure         int64
	Humidity         int64
	Precipitation    float64
	WindSpeed        float64
	AvgWindSpeed     float64
	WindDirection    int64
	AvgWindDirection int64
	CloudCoverage    int64
	Visibility       int64
	AirQuality       int64
}

type OpenWeatherResp struct {
	Code       string               `json:"cod"`
	Message    int                  `json:"message"`
	Count      int                  `json:"cnt"`
	MetricList []OpenWeatherMetrics `json:"list"`
	City       OpenWeatherCity      `json:"city"`
}

type OpenWeatherMetrics struct {
	DT   int64 `json:"dt"`
	Main struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		TempMin     float64 `json:"temp_min"`
		TempMax     float64 `json:"temp_max"`
		Pressure    int64   `json:"pressure"`
		SeaLevel    int64   `json:"sea_level"`
		GroundLevel int64   `json:"grnd_level"`
		Humidity    int64   `json:"humidity"`
		TempKF      float64 `json:"temp_kf"`
	} `json:"main"`
	Weather []struct {
		ID              int64  `json:"id"`
		MainDescription string `json:"main"`
		SubDescription  string `json:"description"`
		Icon            string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int64 `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int64   `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Visibility int64   `json:"visibility"`
	Pop        float64 `json:"pop"`
	Rain       struct {
		OneHour float64 `json:"1h"`
	} `json:"rain"`
	Sys struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DtTxt string `json:"dt_txt"`
}

type OpenWeatherCity struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Country    string `json:"country"`
	Population int64  `json:"population"`
	TimeZone   int64  `json:"timezone"`
	Sunrise    int64  `json:"sunrise"`
	Sunset     int64  `json:"sunset"`
}

type AQIResp struct {
	Status string `json:"status"`
	Data   struct {
		Aqi          int64 `json:"aqi"`
		Idx          int64 `json:"idx"`
		Attributions []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
			Logo string `json:"logo,omitempty"`
		} `json:"attributions"`
		City struct {
			Geo      []float64 `json:"geo"`
			Name     string    `json:"name"`
			URL      string    `json:"url"`
			Location string    `json:"location"`
		} `json:"city"`
		Dominentpol string `json:"dominentpol"`
		Iaqi        struct {
			Dew struct {
				V int64 `json:"v"`
			} `json:"dew"`
			H struct {
				V float64 `json:"v"`
			} `json:"h"`
			No2 struct {
				V float64 `json:"v"`
			} `json:"no2"`
			O3 struct {
				V float64 `json:"v"`
			} `json:"o3"`
			P struct {
				V float64 `json:"v"`
			} `json:"p"`
			Pm25 struct {
				V int64 `json:"v"`
			} `json:"pm25"`
			T struct {
				V float64 `json:"v"`
			} `json:"t"`
			W struct {
				V float64 `json:"v"`
			} `json:"w"`
			Wg struct {
				V float64 `json:"v"`
			} `json:"wg"`
		} `json:"iaqi"`
		Time struct {
			S   string `json:"s"`
			Tz  string `json:"tz"`
			V   int64  `json:"v"`
			Iso string `json:"iso"`
		} `json:"time"`
		Forecast struct {
			Daily struct {
				O3 []struct {
					Avg int64  `json:"avg"`
					Day string `json:"day"`
					Max int64  `json:"max"`
					Min int64  `json:"min"`
				} `json:"o3"`
				Pm10 []struct {
					Avg int64  `json:"avg"`
					Day string `json:"day"`
					Max int64  `json:"max"`
					Min int64  `json:"min"`
				} `json:"pm10"`
				Pm25 []struct {
					Avg int64  `json:"avg"`
					Day string `json:"day"`
					Max int64  `json:"max"`
					Min int64  `json:"min"`
				} `json:"pm25"`
			} `json:"daily"`
		} `json:"forecast"`
		Debug struct {
			Sync time.Time `json:"sync"`
		} `json:"debug"`
	} `json:"data"`
}

func QueryAPI(param map[string]string, method string, url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	// appending to existing query args
	for key, item := range param {
		q.Add(key, item)
	}

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBody, err
}

func GetMetrics(lat float64, lon float64, mType string) (Metrics, error) {
	param := make(map[string]string)
	param["lat"] = fmt.Sprintf("%f", lat)
	param["lon"] = fmt.Sprintf("%f", lon)
	param["appid"] = "b51325554e1adbdfb37ec4cbed1dcfd5"
	param["units"] = "metric"
	url := "https://pro.openweathermap.org/data/2.5/forecast/hourly"

	responseBody, err := QueryAPI(param, http.MethodGet, url)
	if err != nil {
		fmt.Println("QueryAPI error.")
	}

	var openWeatherResp OpenWeatherResp
	if err := json.Unmarshal(responseBody, &openWeatherResp); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	var metric Metrics
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

	responseBody, err = QueryAPI(param, http.MethodGet, url)
	if err != nil {
		fmt.Println("QueryAPI error.")
	}

	var aqiResp AQIResp
	if err := json.Unmarshal(responseBody, &aqiResp); err != nil {
		log.Printf("Unmarshal err, err message: %s", err)
	}

	metric.AirQuality = aqiResp.Data.Aqi

	return metric, nil
}

func GetQuality(metric Metrics) float64 {
	total := RateCloudCoverage(metric.CloudCoverage) +
		RateHumidity(metric.Humidity) +
		RateAvgWindSpeed(int64(metric.AvgWindSpeed)) +
		RateWindDirectionChange(metric.WindDirection, metric.AvgWindDirection) +
		RateAirQuality(metric.AirQuality)
	if total > 0 {
		return float64(total) / 20.00
	} else {
		return 0
	}
}

func RateCloudCoverage(n int64) int64 {
	if n < 30 || n > 70 {
		return 0
	}

	if n < 35 || n > 65 {
		return 1
	}

	if n < 40 || n > 60 {
		return 2
	}

	if n < 45 || n > 55 {
		return 3
	}

	if n > 45 && n < 55 {
		return 4
	}

	return 0
}

func RateHumidity(n int64) int64 {
	if n > 90 {
		return -5
	}

	if n > 70 {
		return -2
	}

	if n > 50 {
		return 0
	}

	if n > 30 {
		return 1
	}

	if n > 10 {
		return 2
	}

	return 0
}

func RateWindDirectionChange(n int64, avg int64) int64 {
	diff := absInt(n - avg)
	if diff < 20 {
		return 0
	}

	if diff < 50 {
		return 1
	}

	if diff < 90 {
		return 2
	}

	if diff < 120 {
		return 3
	}

	if diff > 120 {
		return 4
	}

	return 0
}

func RateAvgWindSpeed(n int64) int64 {
	if n > 13 {
		return 0
	}

	if n > 9 {
		return 1
	}

	if n > 6 {
		return 2
	}

	if n > 4 {
		return 3
	}

	if n > 2 {
		return 4
	}

	return 0
}

func RateAirQuality(n int64) int64 {
	if n > 200 {
		return -10
	}

	if n > 100 {
		return -5
	}

	if n > 50 {
		return -1
	}

	if n > 20 {
		return 1
	}

	if n > 10 {
		return 2
	}

	if n > 0 {
		return 3
	}

	return 0
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
