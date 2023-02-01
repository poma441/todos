package handler

import (
	"net/http"
	"strconv"
	"todos/internal/entity"

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
		return
	}

	c.JSON(http.StatusOK, toDoItemsList)
}

func (h *Handler) AddToDoItem(c *gin.Context) {
	var toDoItemForAdd entity.ToDoItem

	err := c.BindJSON(&toDoItemForAdd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toDoItem, err := h.services.ToDoItem.AddToDoItem(toDoItemForAdd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toDoItem)
}

func (h *Handler) UpdateToDoItem(c *gin.Context) {
	var toDoItemForUpdate entity.ToDoItem

	if err := c.BindJSON(&toDoItemForUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.ToDoItem.UpdateToDoItem(toDoItemForUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Updated successfully"})
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
		return
	}

	c.JSON(http.StatusOK, deletedToDoItemId)
}
