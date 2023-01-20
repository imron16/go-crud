package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUpRouter(app *gin.Engine, handler *handler) {
	apiTermCondition := app.Group("/termCondition")

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang di training golang mysql")
	})

	apiTermCondition.GET("/getAllData", handler.termConditionHandler.GetAllData())
	apiTermCondition.POST("/create", handler.termConditionHandler.CreateData())
	apiTermCondition.POST("/getTermsById", handler.termConditionHandler.GetDataById())
	apiTermCondition.POST("updateTermsCondition", handler.termConditionHandler.UpdateDataById())
	apiTermCondition.POST("/delete", handler.termConditionHandler.DeleteData())

}
