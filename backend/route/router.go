package route

import (
	"github.com/PeterHW963/CVWO/backend/controller"
	"github.com/PeterHW963/CVWO/backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	// MOVE STUFF AROUND FOR ADMINS AND USERS
	router.POST("/authenticate", middleware.Authentication, controller.Authenticate)

	middlewareAdmin := router.Group("/admin", middleware.Authentication, middleware.AdminAuthentication)
	middlewareGeneral := router.Group("/general", middleware.Authentication)

	// middlewareUser:= router.Group("/user", middleware.Authentication, middleware.UserAuthentication)

	//threads
	middlewareAdmin.DELETE("/thread/delete", controller.DeleteThread)
	middlewareGeneral.POST("/thread/create", controller.CreateThread)
	middlewareGeneral.PUT("/thread/update", controller.UpdateThread)
	// router.POST("/thread/create", middleware.Authentication, controller.CreateThread)
	// router.PUT("/thread/update", controller.UpdateThread)
	// router.DELETE("/thread/delete", controller.DeleteThread)
	router.GET("/thread/get", controller.GetThread)
	//posts
	middlewareGeneral.POST("/post/create", controller.CreatePost)
	middlewareGeneral.PUT("/post/update", controller.UpdatePost)
	middlewareAdmin.DELETE("/post/delete", controller.DeletePost)
	router.GET("/post/get", controller.GetPost)
	// router.POST("/post/create", controller.CreatePost)
	// router.PUT("/post/update", controller.UpdatePost)
	// router.DELETE("/post/delete", controller.DeletePost)
	//tags
	middlewareGeneral.POST("/tag/create", controller.CreateTag)
	middlewareGeneral.PUT("/tag/update", controller.UpdateTag)
	middlewareAdmin.DELETE("/tag/delete", controller.DeleteTag)
	router.GET("/tag/get", controller.GetTag)
	// router.POST("/tag/create", controller.CreateTag)
	// router.PUT("/tag/update", controller.UpdateTag)
	// router.DELETE("/tag/delete", controller.DeleteTag)
	//votes
	router.GET("/likedislikepost/get", controller.GetTotalVotesPost)
	router.GET("/likedislikecomment/get", controller.GetTotalVotesComment)
	middlewareGeneral.PUT("/likedislikepost/addupvote", controller.AddDownvotePost)
	middlewareGeneral.PUT("/likedislikecomment/addupvote", controller.AddUpvoteComment)
	middlewareGeneral.PUT("/likedislikepost/adddownvote", controller.AddUpvotePost)
	middlewareGeneral.PUT("/likedislikecomment/adddownvote", controller.AddDownvoteComment)
	//comments
	middlewareGeneral.POST("/comment/create", controller.CreateComment)
	middlewareGeneral.PUT("/comment/update", controller.UpdateComment)
	middlewareAdmin.DELETE("/comment/delete", controller.DeleteComment)
	//users
	router.POST("/user/create", controller.CreateUser)
	middlewareGeneral.PUT("/user/update", controller.UpdateUser)
	middlewareAdmin.DELETE("/user/delete", controller.DeleteUser)
	router.GET("/user/get", controller.GetUser)
	router.POST("/user/login", controller.Login)
}
