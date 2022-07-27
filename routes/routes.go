package routes

import (
	"github.com/gin-gonic/gin"
	"todo_list/api"
	"todo_list/middleware"
)
import "github.com/gin-contrib/sessions"
import "github.com/gin-contrib/sessions/cookie"

func NewRoute() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:tid", api.ShowTask)
			authed.GET("tasks/:uid", api.ListTasks)
			authed.PUT("updatetask/:tid", api.UpdateTask)
			authed.POST("search", api.SearchTask)
			authed.DELETE("task/:tid", api.DeleteTask)
		}
	}
	return r
}
