package service

import (
	"album-manager/album-manager/domain"
	"album-manager/album-manager/models"
	"album-manager/album-manager/utils/date_utils"
	"album-manager/album-manager/utils/errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Service ...
type Service interface {
	InsertImageInAlbum(user models.Image) (models.Image, *errors.RestErr)
	CreateAlbum(album models.Album) (string, *errors.RestErr)
	DeleteAlbum(email, albumName string) (string, *errors.RestErr)
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

func (s *service) InsertImageInAlbum(user models.Image) (models.Image, *errors.RestErr) {
	fmt.Println("implement service")
	// accessTokenID = strings.TrimSpace(accessTokenID)
	// if len(accessTokenID) == 0 {
	// 	return nil, errors.NewBadRequestError("invalid access token id")
	// }
	s.repository.InsertImageInAlbum(&user)
	// accessToken, err := s.repository.GetByID(accessTokenID)
	// if err != nil {
	// 	return nil, err
	// }
	return user, nil
}

func (s *service) CreateAlbum(album models.Album) (string, *errors.RestErr) {
	album.ID = primitive.NewObjectID()
	album.Info.CreatedDate = date_utils.GetNowDBFormat()
	if err := s.repository.CreateAlbum(&album); err != nil {
		return "album creation failed", err
	}
	return "album with album name " + album.AlbumName + " created successfully", nil
}

func (s *service) DeleteAlbum(email, albumName string) (string, *errors.RestErr) {
	if err := s.repository.DeleteAlbum(email + "-" + albumName); err != nil {
		return "album creation failed", err
	}
	return "album with album name " + albumName + " deleted successfully", nil
}
