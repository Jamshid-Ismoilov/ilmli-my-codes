package controllers

import (
	"app/modules/tools"
	"app/modules/verify/base"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var body base.SignupBody

	err := c.BindJSON(&body)

	if err != nil {
		log.Printf("error in binding json: %v", err)
	}

	isExists := base.CheckMail(body)

	if isExists {
		c.JSON(405, "Your email had already been signed up")
	}

	if isExists == false {

		password := tools.SendMail(body.Email)

		base.SetPassword(body, password)
		fmt.Println(password)

		c.JSON(201, "Your password has been sent to your email")

	}
}
