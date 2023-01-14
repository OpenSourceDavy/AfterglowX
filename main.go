package main

import (
	"fmt"
	"net/http"

	"github.com/wwkeyboard/sunsetwx/route"
	"github.com/wwkeyboard/sunsetwx/settings"
)

func main() {
	router := route.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
