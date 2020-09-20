package service

import (
	"album-manager/album-manager/domain"
	"album-manager/album-manager/models"
	"album-manager/album-manager/utils/errors"
	"fmt"
)

//Service ...
type Service interface {
	InsertImageInAlbum(user models.User) (models.User, *errors.RestErr)
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

func (s *service) InsertImageInAlbum(user models.User) (models.User, *errors.RestErr) {
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
