package controllers

import (
	"app/modules/post/base"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("postID"))

	if err != nil {
		log.Printf("failed to convert string to int(strconv) in user module: %v", err)
	}
	user := base.GetPostById(id)

	ctx.JSON(200, user)
}
