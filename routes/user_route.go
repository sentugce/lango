package routes

import (
	"tuce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/god", controllers.GetGod())
	router.POST("/god", controllers.CreateGod())
	//router.POST("/god", controllers.CreateGod)
	/* 	router.GET("/user/:userId", controllers.GetAUser())
	   	router.PUT("/user/:userId", controllers.EditAUser())
	   	router.DELETE("/user/:userId", controllers.DeleteAUser())
	   	router.GET("/users", controllers.GetAllUsers()) */
}
