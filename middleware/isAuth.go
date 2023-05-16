package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DevIdol/Golang_JWT/config"
	"github.com/DevIdol/Golang_JWT/helpers"
	"github.com/DevIdol/Golang_JWT/respository"
	"github.com/DevIdol/Golang_JWT/utils"
	"github.com/gin-gonic/gin"
)

func IsAuth(userRespository respository.UserRespository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizeHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizeHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "your are not logged in"})
			return
		}

		config, _ := config.LoadENV(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		}
		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helpers.ErrorPanic(err_id)
		result, err := userRespository.FindById(id)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "the user beloging to this token no longger exist"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Next()
	}
}
