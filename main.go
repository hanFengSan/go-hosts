package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hanFengSan/go-hosts/handler"
)

var isDebugMode = flag.Bool("debug", false, "debug mode")

func main() {
	log.Println("Starting Server")
	// debug模式设置
	flag.Parse()
	if *isDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 路由设置
	app := gin.Default()
	app.StaticFile("/", "./asset/index.html")
	app.GET("/hosts/discord", handler.GetDiscordHosts)
	app.Run(":9092")
}
