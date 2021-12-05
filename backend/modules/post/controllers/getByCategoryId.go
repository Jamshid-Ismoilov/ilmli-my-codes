package controllers

import (
	"app/modules/post/base"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByCategoryId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("categoryID"))

	if err != nil {
		log.Printf("failed to convert string to int(strconv) in posts module: %v", err)
	}
	posts := base.GetByCategoryIdDB(id)

	ctx.JSON(200, posts)
}
