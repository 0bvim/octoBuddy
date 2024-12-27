package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	// TODO: setup trusted proxies
	initializeRoutes(router)

	err := router.Run(":8042")
	if err != nil {
		fmt.Printf("error starting server: %v\n", err)
		return
	}
}
