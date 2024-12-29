package handler

import "github.com/gin-gonic/gin"

func GetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get user passed as parameter",
	})
}

func GetUserFollowers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get followers from user",
	})
}

func GetUserFollowing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get following from user",
	})
}

func GetStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on get status",
	})
}
