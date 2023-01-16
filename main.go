package main

import (
	"fmt"
	"net/http"

	"github.com/wwkeyboard/sunsetwx/route"
	routeV1 "github.com/wwkeyboard/sunsetwx/route"
	"github.com/wwkeyboard/sunsetwx/settings"
)

func main() {

	router := route.InitRouter()

	routerV1 := router.Group("v1")

	routeV1.Setup(routerV1)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
