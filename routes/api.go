package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rifalafandi314/monitoring-server/controllers"
)

func InitRoute(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/system", controllers.GetDataSystem)
	}
}
