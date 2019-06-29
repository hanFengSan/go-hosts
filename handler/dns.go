package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanFengSan/go-hosts/service"
)

// GetDiscordHosts 获取discord的hosts文件
func GetDiscordHosts(ctx *gin.Context) {
	filePath, err := service.GetSubjectHosts("discord")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, gin.H{"status": 500, "message": err.Error()})
	}
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename="+"discord.hosts.txt")
	ctx.Writer.Header().Set("Content-Type", "text/plain")
	http.ServeFile(ctx.Writer, ctx.Request, filePath)
}
