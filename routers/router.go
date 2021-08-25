package routers

import (
	"go-gin-eg/pkg/setting"
	"go-gin-eg/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建不带中间件的路由
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 标签路由
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tag", v1.CreateTag)
		apiv1.PUT("/tag/:id", v1.UpdateTag)
		apiv1.DELETE("/tag/:id", v1.DeleteTag)

		// 文章路由
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/article/:id", v1.GetArticle)
		apiv1.POST("/article", v1.CreateArticle)
		apiv1.PUT("/article/:id", v1.UpdateArticle)
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
