package service

import (
	"album-manager/album-manager/domain"
	"album-manager/album-manager/models"
	"album-manager/album-manager/utils/date_utils"
	"album-manager/album-manager/utils/errors"

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
	return image, nil
}

func (s *service) CreateAlbum(album models.Album) (models.Album, *errors.RestErr) {
	album.ID = primitive.NewObjectID()
	album.Info.CreatedDate = date_utils.GetNowDBFormat()
	if err := s.repository.CreateAlbum(&album); err != nil {
		return models.Album{}, err
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
	return image, nil
}
