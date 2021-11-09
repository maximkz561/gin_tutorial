package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maximkz561/gin_tutorial/controller"
	"github.com/maximkz561/gin_tutorial/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	r := gin.Default()
	r.GET("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.FindAll())
	})
	r.POST("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.Save(c))
	})
	r.Run()
}
