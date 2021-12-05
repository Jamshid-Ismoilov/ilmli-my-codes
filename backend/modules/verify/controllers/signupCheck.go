package controllers

import (
	"app/modules/tools"
	"app/modules/verify/base"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func SignupCheck(c *gin.Context) {
	var body base.LoginBody
	var id int

	err := c.BindJSON(&body)

	if err != nil {
		log.Printf("error in binding json: %v", err)
	}

	ok := base.CheckSignupPassword(body)

	fmt.Println("this is ok", ok)

	if ok {
		id = base.InsertUserDB(body)

		base.DeletePassword(body)

		c.JSON(201, tools.GenerateJWT(tools.JWTUser{ID: id}))

	} else {

		c.JSON(404, nil)
	}
}
