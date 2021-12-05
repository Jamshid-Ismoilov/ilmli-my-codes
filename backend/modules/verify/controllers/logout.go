package controllers

import (
	"app/modules/verify/base"
	"log"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	var body base.LoginBody

	err := c.BindJSON(&body)

	if err != nil {
		log.Printf("error in binding json: %v", err)
	}

	id := base.CheckUserDB(body)

	if id != 0 {

		base.DeleteUser(body)

		c.JSON(201, "You have logged out")
	} else {
		c.JSON(404, "Password incorrect")
	}

}
