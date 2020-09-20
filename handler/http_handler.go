package handler

import (
	"album-manager/album-manager/models"
	"album-manager/album-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AccessTokenHandler ...
type AlbumHandlerInterface interface {
	InsertImageInAlbum(c *gin.Context)
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

func (handler albumHandler) InsertImageInAlbum(c *gin.Context) {
	// accessTokenID := c.Param("access_token_id")
	// accessToken, err := handler.service.GetByID(accessTokenID)
	handler.service.InsertImageInAlbum(models.User{})
	// if err != nil {
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	// c.JSON(http.StatusOK, accessToken)
	c.JSON(http.StatusNotImplemented, "implement me now !!!")
}
