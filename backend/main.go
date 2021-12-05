package main

import (
	category "app/modules/category/controllers"
	post "app/modules/post/controllers"
	user "app/modules/user/controllers"
	verify "app/modules/verify/controllers"
	wiki "app/modules/wikipedia"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// PATHWAYS

	// users := r.Group("/users")

	// users.POST("/signup", user.Signup)
	// users.POST("/signup/check", user.SignupCheck)
	// users.POST("/login", user.Login)
	// users.POST("/logout", user.Logout)

	users := r.Group("/users")

	users.GET("/", user.GetMe)
	users.GET("/:userID", user.GetById)
	users.GET("/teacher/:userID", user.GetTeachers)

	users.PUT("/update", user.UpdateUser)

	posts := r.Group("/posts")

	posts.GET("/:postID", post.GetById)
	posts.GET("/user/:userID", post.GetByUserId)
	posts.GET("/category/:categoryID", post.GetByCategoryId)

	posts.POST("/new", post.NewPost)
	posts.PUT("/update", post.UpdatePost)

	categories := r.Group("/categories")

	categories.GET("/:categoryID", category.GetById)
	categories.GET("/total", category.GetAllCategories)

	verifies := r.Group("/verify")

	verifies.POST("/signup", verify.Signup)
	verifies.POST("/signupCheck", verify.SignupCheck)
	verifies.POST("/login", verify.Login)
	verifies.POST("/logout", verify.Logout)

	wikipedia := r.Group("/wikipedia")

	wikipedia.GET("/pages/:names", wiki.GetPages)
	wikipedia.GET("/page/:name", wiki.GetPage)

	r.Run()
}
