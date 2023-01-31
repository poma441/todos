package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetToDoItemsList(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toDoItemsList, err := h.services.ToDoItem.GetToDoItemsList(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, toDoItemsList)
}

func (h *Handler) AddToDoItem(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toDoItem, err := h.services.ToDoItem.AddToDoItem(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, toDoItem)
}

func (h *Handler) UpdateToDoItem(c *gin.Context) {
	toDoItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toDoItem, err := h.services.ToDoItem.UpdateToDoItem(toDoItemId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, toDoItem)
}

func (h *Handler) DeleteToDoItem(c *gin.Context) {
	toDoItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedToDoItemId, err := h.services.ToDoItem.DeleteToDoItem(toDoItemId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, deletedToDoItemId)
}