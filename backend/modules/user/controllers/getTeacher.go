package controllers

import (
	"app/modules/user/base"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTeachers(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("userID"))

	if err != nil {
		log.Printf("failed to convert string to int(strconv) in word module: %v", err)
	}
	teachers := base.GetTeacherById(id)

	ctx.JSON(200, teachers)
}
