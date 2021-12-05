package controllers

import (
	"app/modules/user/base"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("userID"))

	if err != nil {
		log.Printf("failed to convert string to int(strconv) in user module: %v", err)
	}
	user := base.GetUserById(id)

	ctx.JSON(200, user)
}
