package controllers

import (
	"app/modules/post/base"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByUserId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("userID"))

	if err != nil {
		log.Printf("failed to convert string to int(strconv) in posts module: %v", err)
	}
	posts := base.GetByUserIdDB(id)

	ctx.JSON(200, posts)
}
