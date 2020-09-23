package handler

import (
	"album-manager/album-manager/models"
	"album-manager/album-manager/service"
	"album-manager/album-manager/utils/errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AlbumHandlerInterface ...
type AlbumHandlerInterface interface {
	InsertImageInAlbum(c *gin.Context)
	CreateAlbum(c *gin.Context)
	DeleteAlbum(c *gin.Context)
	DeleteImage(c *gin.Context)
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
	res, err := handler.service.CreateAlbum(album)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, res)

}

//DeleteAlbum ...
func (handler albumHandler) DeleteAlbum(c *gin.Context) {
	album_id := c.Param("album_id")
	res, err := handler.service.DeleteAlbum(album_id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, res)

}

//DeleteImage ...
func (handler albumHandler) DeleteImage(c *gin.Context) {
	albumId := c.Query("album_id")
	imageId := c.Query("image_id")
	res, err := handler.service.DeleteImage(albumId, imageId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (handler albumHandler) InsertImageInAlbum(c *gin.Context) {
	var image models.Image
	_, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	image.Info.CreatedBy = c.PostForm("created_by")
	file, hdr, err := c.Request.FormFile("image_data")
	image.AlbumId = c.PostForm("album_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	image.ImageName = hdr.Filename
	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	image.ImageData = data
	defer file.Close()

	res, errRes := handler.service.InsertImageInAlbum(image)
	if err != nil {
		c.JSON(errRes.Status, err)
		return
	}
	c.JSON(http.StatusOK, res)

}
