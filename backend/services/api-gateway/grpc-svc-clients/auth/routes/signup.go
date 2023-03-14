package routes

import (
	"context"
	"net/http"
	"todos/services/api-gateway/grpc-svc-clients/auth/pb"
	"todos/services/api-gateway/models"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	FullName string `json:"fullname"`
}

func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	requestInfo := &models.RequestAdditionalInfo{
		UserAgent: ctx.Request.UserAgent(),
		SrcIP:     ctx.ClientIP(),
	}

	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	res, _ := c.Register(context.Background(), &pb.RegisterRequest{ // Вызов процедуры сервиса auth
		Email:    body.Email,
		Password: body.Password,
		Role:     body.Role,
		Phone:    body.Phone,
		Fullname: body.FullName,
		RequestInfo: &pb.RequestAdditionalInfo{
			UserAgent: requestInfo.UserAgent,
			SrcIpAddr: requestInfo.SrcIP,
		},
	})

	if !res.GetSuccess() {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"success": res.GetSuccess(), "error": res.GetError()})
		return
	}

	ctx.SetCookie("refresh_token", res.GetTokensInfo().GetRefreshToken(), 0, res.GetTokensInfo().GetRefreshCookiePath(), res.GetTokensInfo().GetRefreshCookieHost(), false, true)
	ctx.SetCookie("logout_token", res.GetTokensInfo().GetRefreshToken(), 0, res.GetTokensInfo().GetLogoutCookiePath(), res.GetTokensInfo().GetLogoutCookieHost(), false, true)
	ctx.JSON(http.StatusOK, gin.H{"success": res.GetSuccess(), "access_token": res.GetTokensInfo().GetAccessToken()})
}
