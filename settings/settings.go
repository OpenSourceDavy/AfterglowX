package settings

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/wwkeyboard/sunsetwx/logs"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize            int
	JwtSecret           string
	JwtExpHour          int
	RefreshTokenSecret  string
	RefreshTokenExpHour int

	WindyKey string
	WindyURL string

	OpenWeatherAppID string
	OpenWeatherURL   string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logs.Log.Error("Fail to parse 'conf/app.ini': %v", err)

	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadAPIKeys()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logs.Log.Error("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logs.Log.Error("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	JwtExpHour = sec.Key("JWT_SECRET_EXPIRATION_HOUR").MustInt(2)
	RefreshTokenSecret = sec.Key("REFRESH_TOKEN_SECRET").MustString("!@)*#)!@U#@*!(!)")
	RefreshTokenExpHour = sec.Key("REFRESH_TOKEN_EXPIRATION_HOUR").MustInt(2)
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadAPIKeys() {
	sec, err := Cfg.GetSection("api")
	if err != nil {
		logs.Log.Error("Fail to get section 'api': %v", err)
	}
	WindyKey = sec.Key("WINDY_KEY").String()
	WindyURL = sec.Key("WINDY_URL").String()
	OpenWeatherAppID = sec.Key("OPEN_WEATHER_APP_ID").String()
	OpenWeatherURL = sec.Key("OPEN_WEATHER_URL").String()
}
