package controllers

import (
	"app/modules/post/base"
	"app/modules/tools"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdatePost(ctx *gin.Context) {
	var body base.UpdatePostBody

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

	post := base.UpdatePost(body, jwt.ID)

	ctx.JSON(201, post)
}
