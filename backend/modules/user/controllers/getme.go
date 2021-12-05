package controllers

import (
	"app/modules/tools"
	"app/modules/user/base"
	"log"

	"github.com/gin-gonic/gin"
)

func GetMe(ctx *gin.Context) {

	token := ctx.Request.Header.Get("param")

	jwt, err := tools.ParseJWT(token)

	if err != nil {
		log.Printf("jwt error: %v", err)
	}
	user := base.GetUserById(jwt.ID)

	ctx.JSON(200, user)
}
