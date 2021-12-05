package controllers

import (
	"app/modules/tools"
	"app/modules/verify/base"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body base.LoginBody

	err := c.BindJSON(&body)

	if err != nil {
		log.Printf("error in binding json: %v", err)
	}

	id := base.CheckUserDB(body)

	if id != 0 {
		c.JSON(201, tools.GenerateJWT(tools.JWTUser{ID: id}))

	} else {
		c.JSON(404, nil)
	}

}
