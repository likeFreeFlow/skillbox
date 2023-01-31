package handler

import (
	gin "github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	users := router.Group("/user")
	{
		users.POST("/create", h.createUser)
		users.POST("/addFriend", h.addFriend)
		users.DELETE(":id", h.deleteUser)
		users.GET(":id", h.showFriends)
		users.PUT(":id", h.setAge)

	}
	return router
}
