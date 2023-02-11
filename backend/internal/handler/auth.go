package handler

import (
	"log"
	"net/http"
	"todos/config"
	"todos/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Signup(c *gin.Context) {
	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Password = h.services.Authorization.HashPass(input.Password)

	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (h *Handler) Signin(c *gin.Context) {
	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.Authorization.GetUser(input.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Authorization.ComparePass(user.Password, input.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cfg, err := config.InitConfig("../../config")
	if err != nil {
		log.Fatal("Ошибка инициализации ", err)
	}

	access_token, err := h.services.CreateToken(cfg.Access.Ttl, user.Id, cfg.Access.PrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, cfg.Access.MaxAge*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
