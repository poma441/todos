package handler

import (
	"log"
	"net/http"
	"todos/config"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
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

func (h *Handler) ValidateUser(c *gin.Context) {
	var input map[string]interface{}
	var schemaLoader gojsonschema.JSONLoader
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch len(input) {
	case 2:
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/murex/go/src/todos/backend/internal/form/user.json")
	case 5:
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/murex/go/src/todos/backend/internal/form/schema.json")
	default:
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/murex/go/src/todos/backend/internal/form/user.json")
	}

	loader := gojsonschema.NewGoLoader(input)

	result, err := gojsonschema.Validate(schemaLoader, loader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed"})
		return
	}

}
