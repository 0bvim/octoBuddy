package handler

import "github.com/gin-gonic/gin"

func GetMe(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get Me",
	})
}

func ListFollowers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on list followers",
	})
}

func ListFollowing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on list following",
	})
}

func GetAllowList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on get allow list",
	})
}

func AddAllowList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on add allow list",
	})
}

func DeleteAllowList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on delete allow list",
	})
}

func GetDenyList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on get deny list",
	})
}

func AddDenyList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on add deny list",
	})
}

func DeleteDenyList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on delete deny list",
	})
}
