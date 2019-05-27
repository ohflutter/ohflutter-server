package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ohflutter-server/pkg/app"
	"ohflutter-server/pkg/e"
	"ohflutter-server/services/user_service"
)

func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}

	userService := user_service.User{
		UserID: 1,
		Mobile: "1858138074",
		Name:   "Jack",
		City:   "成都",
	}

	if err := userService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAILED, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
