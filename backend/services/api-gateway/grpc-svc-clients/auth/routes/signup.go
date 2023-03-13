package routes

import (
	"context"
	"net/http"
	"todos/internal/entity"
	"todos/services/api-gateway/grpc-svc-clients/auth/pb"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"passsword"`
}

func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	// Формирование информации о запросе
	requestInfo := &entity.RequestAdditionalInfo{
		UserAgent: ctx.Request.UserAgent(),
		SrcIP:     ctx.ClientIP(),
	}

	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
		RequestInfo: &pb.RequestAdditionalInfo{
			UserAgent: requestInfo.UserAgent,
			SrcIpAddr: requestInfo.SrcIP,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
