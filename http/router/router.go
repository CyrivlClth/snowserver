package router

import (
	"github.com/gin-gonic/gin"

	"github.com/CyrivlClth/snowserver/http/api"
)

func New() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", api.NextID)
	router.GET("/stats", api.Stats)
	router.GET("/count/:count", api.GetIDs)

	return
}
