package main

import (
	"os"

	"github.com/rifalafandi314/monitoring-server/engines"
)

func main(){
	r := engines.SetEngine()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port="8000"
	}

	
	r.Run("0.0.0.0:" + port)
}
