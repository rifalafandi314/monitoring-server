package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifalafandi314/monitoring-server/services"
)


func GetDataSystem(ctx *gin.Context) {
	data := services.GetSystemStat()
	ctx.JSON(http.StatusOK, data)
}