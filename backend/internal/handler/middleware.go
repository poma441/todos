package handler

import (
	"log"
	"net/http"
	"todos/config"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserIdentify(c *gin.Context) {

	access_token, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if access_token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Вы не авторизированы"})
		return
	}

	cfg, err := config.InitConfig("../../config")
	if err != nil {
		log.Fatal("Ошибка инициализации ", err)
	}

	sub, err := h.services.Authorization.ValidateToken(access_token, cfg.Access.PublicKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.services.Authorization.GetUserById(int(sub.(float64)))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Пользователя с таким токеном не существует"})
		return
	}

	c.Set("currentUser", user)
	c.Next()

}
