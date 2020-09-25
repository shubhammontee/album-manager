package service

import (
	"album-manager/album-manager/domain"
	"album-manager/album-manager/models"
	"album-manager/album-manager/utils/date_utils"
	"album-manager/album-manager/utils/errors"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Service ...
type Service interface {
	InsertImageInAlbum(image models.Image) (models.Image, *errors.RestErr)
	CreateAlbum(album models.Album) (models.Album, *errors.RestErr)
	DeleteAlbum(id string) (models.Album, *errors.RestErr)
	DeleteImage(albumId, imageId string) (models.Image, *errors.RestErr)
}

type service struct {
	repository domain.AlbumDaoInterface
}

//NewService ...
func NewService(repo domain.AlbumDaoInterface) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) InsertImageInAlbum(image models.Image) (models.Image, *errors.RestErr) {
	image.ID = primitive.NewObjectID()
	image.Info.CreatedDate = date_utils.GetNowDBFormat()
	err := s.repository.InsertImageInAlbum(&image)
	if err != nil {
		return models.Image{}, err
	}

	restErr := sendNotification(fmt.Sprintf("inserted image with name %s with id of %s", image.ImageName, image.FileId))
	if restErr != nil {
		return models.Image{}, nil
	}
	return image, nil
}

func (s *service) CreateAlbum(album models.Album) (models.Album, *errors.RestErr) {
	album.ID = primitive.NewObjectID()
	album.Info.CreatedDate = date_utils.GetNowDBFormat()
	if err := s.repository.CreateAlbum(&album); err != nil {
		return models.Album{}, err
	}
	err := sendNotification(fmt.Sprintf("new album with album name %s created", album.AlbumName))
	if err != nil {
		return models.Album{}, nil
	}
	return album, nil
}

func (s *service) DeleteAlbum(id string) (models.Album, *errors.RestErr) {
	var album models.Album
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Album{}, errors.NewInternalServerError(err.Error())
	}
	album.ID = oId
	if err := s.repository.DeleteAlbum(&album); err != nil {
		return models.Album{}, err
	}
	restErr := sendNotification(fmt.Sprintf("album with album name %s deleted", album.AlbumName))
	if restErr != nil {
		return models.Album{}, nil
	}
	return album, nil
}

func (s *service) DeleteImage(albumId, imageId string) (models.Image, *errors.RestErr) {
	var image models.Image
	oId, err := primitive.ObjectIDFromHex(imageId)
	if err != nil {
		return models.Image{}, errors.NewInternalServerError(err.Error())
	}
	image.ID = oId
	image.AlbumId = albumId
	if err := s.repository.DeleteImage(&image); err != nil {
		return models.Image{}, err
	}
	restErr := sendNotification(fmt.Sprintf("image with name %s and id of %s deleted", image.ImageName, image.FileId))
	if restErr != nil {
		return models.Image{}, nil
	}
	return image, nil
}

func sendNotification(message string) *errors.RestErr {
	req := &struct {
		Message string `form:"message" json:"message"`
	}{Message: message}

	log.Printf("Write file to DB was successful. before post File size:")
	jsonReq, restErr := json.Marshal(req)
	if restErr != nil {
		return errors.NewInternalServerError(restErr.Error())
	}

	_, restErr = http.Post("http://host.docker.internal:9000/sendnotification", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if restErr != nil {
		return errors.NewInternalServerError(restErr.Error())
	}
	log.Printf("after post")
	return nil

}
