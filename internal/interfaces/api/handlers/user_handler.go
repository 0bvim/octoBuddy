package handlers

import (
	"net/http"

	"github.com/0bvim/octoBuddy/internal/application/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var userIDString string
	if userID, ok := userID.(string); ok {
		userIDString = userID
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID is not a string"})
		return
	}

	user, err := h.userService.GetUser(userIDString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}

	// token := c.GetHeader("Authorization")
	// followers, err := h.userService.FetchFollowers(token)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user followers"})
	// 	return
	// }
	// fmt.Println(user.Followed)
	// fmt.Println(user.Follower)

	// //FIXME:
	// user.Follower = append(user.Follower, followers...)
	// user.Followed = followers
	c.JSON(http.StatusOK, user)
}
