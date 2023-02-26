package handler

import (
	"net/http"
	"strings"
	"todos/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

func (h *Handler) UserIdentify(c *gin.Context) {
	// Формирование информации о запросе
	requestInfo := &entity.RequestAdditionalInfo{
		UserAgent: c.Request.UserAgent(),
		IP:        c.ClientIP(),
	}

	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	access_token := headerParts[1]
	if len(access_token) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := h.services.Authorization.ValidateToken(access_token, h.config.Token.Keys.PublicKey, false, requestInfo)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ??? Подумать, возможно, не лучшее решение
	// user, err := h.services.Authorization.GetUserById(int(userId.(float64)))
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	c.Set("CurrentUserId", userId)
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
