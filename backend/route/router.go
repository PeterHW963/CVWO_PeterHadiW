package route

import (
	"github.com/PeterHW963/CVWO/backend/controller"
	// "github.com/PeterHW963/CVWO/backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	// (might not be needed) middlewares
	// router.POST("/authenticate", middleware.Authentication, controller.Authenticate)
	// middlewareAdmin := router.Group("/admin", middleware.Authentication, middleware.AdminAuthentication)
	// middlewareGeneral := router.Group("/general", middleware.Authentication)
	// middlewareUser := router.Group("/user", middleware.Authentication, middleware.UserAuthentication)

	//threads
	router.DELETE("/thread/delete", controller.DeleteThread)
	router.POST("/thread/create", controller.CreateThread)
	router.PUT("/thread/update", controller.UpdateThread)
	router.GET("/thread/get", controller.GetThread)
	//posts
	router.POST("/post/create", controller.CreatePost)
	router.PUT("/post/update", controller.UpdatePost)
	router.DELETE("/post/delete", controller.DeletePost)
	router.GET("/post/get", controller.GetPost)
	//tags
	router.POST("/tag/create", controller.CreateTag)
	router.PUT("/tag/update", controller.UpdateTag)
	router.DELETE("/tag/delete", controller.DeleteTag)
	router.GET("/tag/get", controller.GetTag)

	//votes
	router.GET("/likedislikepost/get", controller.GetTotalVotesPost)
	router.GET("/likedislikecomment/get", controller.GetTotalVotesComment)
	router.PUT("/likedislikepost/addupvote", controller.AddDownvotePost)
	router.PUT("/likedislikecomment/addupvote", controller.AddUpvoteComment)
	router.PUT("/likedislikepost/adddownvote", controller.AddUpvotePost)
	router.PUT("/likedislikecomment/adddownvote", controller.AddDownvoteComment)
	//comments
	router.POST("/comment/create", controller.CreateComment)
	router.PUT("/comment/update", controller.UpdateComment)
	router.DELETE("/comment/delete", controller.DeleteComment)
	//users
	router.POST("/user/create", controller.CreateUser)
	router.PUT("/user/update", controller.UpdateUser)
	router.DELETE("/user/delete", controller.DeleteUser)
	router.GET("/user/get", controller.GetUser)
	router.POST("/user/login", controller.Login)
}
