package core

import (
	"github.com/gin-gonic/gin"
	"ipData/routes"
	"log"
	"net/http"
)

func StartApp() {
	// app启动方法

	gin.SetMode(RunTimeMode)

	server := &http.Server{
		Addr:           Addr,
		Handler:        routes.InitRouter(),
		ReadTimeout:    READ_TIMEOUT,
		WriteTimeout:   WRITE_TIMEOUT,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
