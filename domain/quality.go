package domain

import "time"

type QualityRequest struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Type string  `json:"type"`
}

type QualityResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

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

type QualityUseCase interface {
	GetSunsetQuality(lat float64, lon float64) (float64, error)
	GetSunriseQuality(lat float64, lon float64) (float64, error)
}
