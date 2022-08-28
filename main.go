package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-shop/docs"
	"go-shop/utils/graceful"
	"log"
	"net/http"
	"time"
)

func registerMiddleware(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(params gin.LogFormatterParams) string {
				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s %d %s %s\"\n",
					params.ClientIP,
					params.TimeStamp.Format(time.RFC3339),
					params.Method,
					params.Path,
					params.Request.Proto,
					params.StatusCode,
					params.Latency,
					params.ErrorMessage,
				)
			},
		),
	)
	r.Use(gin.Recovery())
}

// @title 电商项目
// @description 电商项目
// @version 1.0
// @contact.name go-shop
// @contact.url https://www.go-shop.com
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	registerMiddleware(r)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s", err)
		}
	}()

	graceful.ShutdownGin(srv, time.Second*5)

}
