package domain

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

type WindyResp struct {
	Ts    []int64 `json:"ts"`
	Units struct {
		TempSurface     string `json:"temp-surface"`
		DewpointSurface string `json:"dewpoint-surface"`
		LcloudsSurface  string `json:"lclouds-surface"`
		McloudsSurface  string `json:"mclouds-surface"`
		HcloudsSurface  string `json:"hclouds-surface"`
	} `json:"units"`
	TempSurface     []float64 `json:"temp-surface"`
	DewpointSurface []float64 `json:"dewpoint-surface"`
	LcloudsSurface  []float64 `json:"lclouds-surface"`
	McloudsSurface  []float64 `json:"mclouds-surface"`
	HcloudsSurface  []float64 `json:"hclouds-surface"`
	Warning         string    `json:"warning"`
}

type QualityUseCase interface {
	GetSunsetQuality(lat float64, lon float64) (float64, error)
	GetSunriseQuality(lat float64, lon float64) (float64, error)
}
