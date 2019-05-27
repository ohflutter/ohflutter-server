package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ohflutter-server/pkg/app"
	"ohflutter-server/pkg/e"
	"ohflutter-server/services/article_service"
)

func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	articleService := article_service.Article{
		UserID: 2,
		Title:  "李琪的第一个文章",
	}

	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAILED, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
