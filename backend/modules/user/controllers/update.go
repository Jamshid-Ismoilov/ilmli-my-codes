package controllers

import (
	"app/modules/tools"
	"app/modules/user/base"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {
	var body base.UpdateUserBody

	err := ctx.BindJSON(&body)

	if err != nil {
		log.Printf("error in binding json: %v", err)
	}

	token := ctx.Request.Header.Get("param")
	fmt.Println(token)

	jwt, err := tools.ParseJWT(token)

	if err != nil {
		ctx.JSON(405, "Cannot upload")
	}

	user := base.UpdateUserDB(body, jwt.ID)

	ctx.JSON(201, user)
}
