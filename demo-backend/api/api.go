package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/v2"
	"github.com/sillyhatxu/gin-utils/v2/entity"
	"github.com/sillyhatxu/gin-utils/v2/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitialAPI(port int) {
	router := SetupRouter()
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Errorf("server down. %v", err)
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	router, err := ginutils.SetupRouter()
	if err != nil {
		panic(err)
	}
	userGroup := router.Group("/demo")
	{
		userGroup.GET("", demo)
	}
	return router
}

func demo(ctx *gin.Context) {
	params := map[string]interface{}{
		"id":        1,
		"user_name": "Sillyhat Xu",
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(params)))
	return
}
