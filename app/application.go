package app

import (
	"album-manager/album-manager/domain"
	"album-manager/album-manager/handler"
	"album-manager/album-manager/service"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	dbRepository := domain.NewDao()
	atService := service.NewService(dbRepository)
	atHandler := handler.NewHandler(atService)
	// router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/insertImageInAlbum", atHandler.InsertImageInAlbum)
	router.POST("/createAlbum", atHandler.CreateAlbum)
	router.DELETE("/deleteAlbum/:album_id", atHandler.DeleteAlbum)
	router.DELETE("/deleteImage", atHandler.DeleteImage)
	router.Run(":8000")

}
