package domain

import (
	"album-manager/album-manager/datasource/mongodb"
	"album-manager/album-manager/models"
	"album-manager/album-manager/utils/errors"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	database *mongo.Database
)

type AlbumDaoInterface interface {
	InsertImageInAlbum(user *models.Image) *errors.RestErr
	CreateAlbum(album *models.Album) *errors.RestErr
	DeleteAlbum(album string) *errors.RestErr
}

type albumDao struct {
}

//NewService ...
func NewDao() AlbumDaoInterface {
	return &albumDao{}

}
func init() {
	database = mongodb.GetMongoInstance().Database("album_manager")
	//collection = mongodb.GetMongoInstance().Database("usersdb").Collection("users")
}

func (ad albumDao) CreateAlbum(album *models.Album) *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := database.CreateCollection(ctx, album.Email+"-"+album.AlbumName)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	collection := mongodb.GetMongoInstance().Database("album_manager").Collection("albums_list")
	_, err = collection.InsertOne(ctx, album)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (ad albumDao) DeleteAlbum(album string) *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := database.Collection(album).Drop(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//Insert user ingto the database
func (ad albumDao) InsertImageInAlbum(user *models.Image) *errors.RestErr {
	fmt.Println("implement dao")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// _, err := collection.InsertOne(ctx, user)
	// if err != nil {
	// 	return errors.NewInternalServerError(err.Error())
	// }
	// id := fmt.Sprintf("%v", res.InsertedID)
	// user.ID = id[10 : len(id)-2]
	return nil
}

//GetUser get single user from users db
// func (user *User) GetUser(id string) *errors.RestErr {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	if err := collection.FindOne(ctx, User{ID: user.ID}).Decode(&user); err != nil {
// 		return errors.NewInternalServerError(err.Error())
// 	}
// 	return nil
// }

// //GetAllUser ...
// func (user *User) GetAllUser() ([]User, *errors.RestErr) {
// 	users := make([]User, 0)
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	cursor, err := collection.Find(ctx, User{})
// 	if err != nil {
// 		return nil, errors.NewInternalServerError(err.Error())
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var user User
// 		cursor.Decode(&user)
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

// //UpdateUser ...
// func (user *User) UpdateUser() *errors.RestErr {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	updateBson := bson.M{}
// 	if user.FirstName != "" {
// 		updateBson["first_name"] = user.FirstName
// 	}
// 	if user.LastName != "" {
// 		updateBson["last_name"] = user.LastName
// 	}
// 	if user.Email != "" {
// 		updateBson["email"] = user.Email
// 	}
// 	if user.Status != "" {
// 		updateBson["status"] = user.Status
// 	}
// 	update := bson.M{"$set": updateBson}
// 	result, err := collection.UpdateOne(ctx, User{ID: user.ID}, update)
// 	if err != nil {
// 		return errors.NewInternalServerError(err.Error())
// 	}
// 	fmt.Println(result.ModifiedCount)
// 	return nil
// }

// //DeleteUser ...
// func (user *User) DeleteUser() *errors.RestErr {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	_, err := collection.DeleteOne(ctx, User{ID: user.ID})
// 	if err != nil {
// 		return errors.NewInternalServerError(err.Error())
// 	}
// 	return nil
// }
