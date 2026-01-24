package engines

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rifalafandi314/monitoring-server/routes"
)


func SetEngine() *gin.Engine {
	r := gin.Default()

	_= godotenv.Load()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5731"},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "APTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "ngrok-skip-browser-warning"},
		AllowCredentials: true,
	}))


	routes.InitRoute(r)
	return r
}