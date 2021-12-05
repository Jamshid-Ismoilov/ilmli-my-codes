package controllers

import (
	"app/modules/category/base"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("categoryID"))

	if err != nil {
		log.Printf("failed to convert string to int(strconv) in category module: %v", err)
	}
	category := base.GetCategoryById(id)

	ctx.JSON(200, category)
}
