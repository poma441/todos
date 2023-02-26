package handler

import (
	"net/http"
	"todos/internal/entity"
	hasher "todos/pkg/password_hasher"

	"github.com/gin-gonic/gin"
)

/*
*	Обработчик запроса регистрации пользователя
 */
func (h *Handler) SignUp(c *gin.Context) {

	// Формирование информации о запросе
	requestInfo := &entity.RequestAdditionalInfo{
		UserAgent: c.Request.UserAgent(),
		IP:        c.ClientIP(),
	}

	var input entity.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Password = hasher.HashPass(input.Password)

	createdUserId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерация access токена
	accessToken, err := h.services.Authorization.CreateToken(createdUserId, h.config.Token.Access.TTL, h.config.Token.Keys.PrivateKey, false, requestInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерация refresh токена
	refreshToken, err := h.services.Authorization.CreateToken(createdUserId, h.config.Token.Refresh.TTL, h.config.Token.Keys.PrivateKey, true, requestInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", refreshToken, 0, h.config.Token.Refresh.RefreshCookiePath, h.config.Server.Host, false, true)
	c.SetCookie("logout_refresh_token", refreshToken, 0, h.config.Token.Refresh.LogoutCookiePath, h.config.Server.Host, false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

/*
*	Обработчик запроса входа пользователя
 */
func (h *Handler) SignIn(c *gin.Context) {
	// Формирование информации о запросе
	requestInfo := &entity.RequestAdditionalInfo{
		UserAgent: c.Request.UserAgent(),
		IP:        c.ClientIP(),
	}

	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.Authorization.GetUser(input.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := hasher.ComparePass(user.Password, input.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Генерация access токена
	accessToken, err := h.services.Authorization.CreateToken(user.Id, h.config.Token.Access.TTL, h.config.Token.Keys.PrivateKey, false, requestInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерация refresh токена
	refreshToken, err := h.services.Authorization.CreateToken(user.Id, h.config.Token.Refresh.TTL, h.config.Token.Keys.PrivateKey, true, requestInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", refreshToken, 0, h.config.Token.Refresh.RefreshCookiePath, h.config.Server.Host, false, true)
	c.SetCookie("logout_refresh_token", refreshToken, 0, h.config.Token.Refresh.LogoutCookiePath, h.config.Server.Host, false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

/*
*	Обработчик запроса обновления пары токенов
 */
func (h *Handler) Refresh(c *gin.Context) {

	// Формирование информации о запросе
	requestInfo := &entity.RequestAdditionalInfo{
		UserAgent: c.Request.UserAgent(),
		IP:        c.ClientIP(),
	}

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Проверяем предоставленный refresh токен на валидность
	userIdClaim, err := h.services.Authorization.ValidateToken(refreshToken, h.config.Token.Keys.PublicKey, true, requestInfo)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Инвалидируем прошлый refresh токен пользователя
	if err := h.services.Authorization.InvalidateRefreshToken(refreshToken, h.config.Token.Keys.PublicKey); err != nil {
		// Логгировать действие, если какая то ошибка, ???????? (пока не придумал)
	}

	userId := int(userIdClaim.(float64))

	// Генерация access токена
	newAccessToken, err := h.services.Authorization.CreateToken(userId, h.config.Token.Access.TTL, h.config.Token.Keys.PrivateKey, false, requestInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерация refresh токена
	newRefreshToken, err := h.services.Authorization.CreateToken(userId, h.config.Token.Refresh.TTL, h.config.Token.Keys.PrivateKey, true, requestInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", newRefreshToken, 0, h.config.Token.Refresh.RefreshCookiePath, h.config.Server.Host, false, true)
	c.SetCookie("logout_refresh_token", newRefreshToken, 0, h.config.Token.Refresh.LogoutCookiePath, h.config.Server.Host, false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": newAccessToken})
}

/*
*	Обработчик logout'a пользователя
 */
func (h *Handler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("logout_refresh_token")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.services.Authorization.InvalidateRefreshToken(refreshToken, h.config.Token.Keys.PublicKey); err != nil {
		// Логгировать действие, если какая то ошибка, все равно будем посылать фронту, чтоб вылогинил юзера
	}

	c.SetCookie("refresh_token", "", -1, h.config.Token.Refresh.RefreshCookiePath, h.config.Server.Host, false, true)
	c.SetCookie("logout_refresh_token", "", -1, h.config.Token.Refresh.LogoutCookiePath, h.config.Server.Host, false, true)
	c.Status(http.StatusOK)
}
