package controllers

import (
	"app/modules/category/base"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(ctx *gin.Context) {

	ctx.JSON(200, base.GetAllCategories())
}
