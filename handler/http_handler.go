package handler

import (
	"album-manager/album-manager/models"
	"album-manager/album-manager/service"
	"album-manager/album-manager/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AlbumHandlerInterface ...
type AlbumHandlerInterface interface {
	InsertImageInAlbum(c *gin.Context)
	CreateAlbum(c *gin.Context)
	DeleteAlbum(c *gin.Context)
}
type albumHandler struct {
	service service.Service
}

//NewHandler ...
func NewHandler(service service.Service) AlbumHandlerInterface {
	return &albumHandler{
		service: service,
	}
}

//CreateAlbum ...
func (handler albumHandler) CreateAlbum(c *gin.Context) {
	//albumName := c.Param("email_id")
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	// accessToken, err := handler.service.GetByID(accessTokenID)
	message, err := handler.service.CreateAlbum(album)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": message})

}

//DeleteAlbum ...
func (handler albumHandler) DeleteAlbum(c *gin.Context) {
	//albumName := c.Param("email_id")
	email := c.Query("email_id")
	albumName := c.Query("album_name")
	message, err := handler.service.DeleteAlbum(email, albumName)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": message})

}

func (handler albumHandler) InsertImageInAlbum(c *gin.Context) {
	// accessTokenID := c.Param("access_token_id")
	// accessToken, err := handler.service.GetByID(accessTokenID)
	handler.service.InsertImageInAlbum(models.Image{})
	// if err != nil {
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	// c.JSON(http.StatusOK, accessToken)
	c.JSON(http.StatusNotImplemented, "implement me now !!!")
}
