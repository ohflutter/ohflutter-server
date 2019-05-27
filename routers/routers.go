package routers

import (
	"github.com/gin-gonic/gin"
	v1 "ohflutter-server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	apiv1.Use()
	{
		//
		apiv1.GET("/users", v1.AddUser)
		apiv1.GET("/articles", v1.AddArticle)
	}

	return r
}
