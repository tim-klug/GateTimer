package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Init() {
	ginRouter := gin.Default()
	ginRouter.Use(Cors())

	v1 := ginRouter.Group("api/v1")
	NewSetGateTime(*v1)
	NewAddGateTime(*v1)
	NewShowGateTimes(*v1)
	NewOpenGate(*v1)

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}

	err := ginRouter.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server is running on port " + port)
}

func OptionsGate(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	context.Next()
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
