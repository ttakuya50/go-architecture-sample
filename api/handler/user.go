package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ttakuya50/go-architecture-sample/api/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type RegisterRequest struct {
	Name string `json:"name"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var json RegisterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Register(c, json.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": json.Name})
}

type DeleteRequest struct {
	UserID int64 `json:"user_id"`
}

func (h *UserHandler) Delete(c *gin.Context) {
	var json DeleteRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Delete(c, json.UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

type AddListRequest struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
}

func (h *UserHandler) AddList(c *gin.Context) {
	var json AddListRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.AddList(c, json.UserID, json.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
