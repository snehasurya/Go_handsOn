package routes

import (
	middlewares "example.com/restApiGin/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	authRoutes := server.Group("/")
	authRoutes.Use(middlewares.AuthenticateUser)
	authRoutes.POST("/events", postEvent)
	authRoutes.PUT("/events/:id", putEvent)
	authRoutes.DELETE("/events/:id", deleteEvent)

	//server.POST("/events", middlewares.AuthenticateUser, postEvent)
	//server.PUT("/events/:id", putEvent)
	//server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", userSignUp)
	server.POST("/login", userLogin)
}
