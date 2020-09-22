package domain

import (
	"album-manager/album-manager/datasource/mongodb"
	"album-manager/album-manager/models"
	"album-manager/album-manager/utils/errors"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var (
	database *mongo.Database
)

type AlbumDaoInterface interface {
	InsertImageInAlbum(image *models.Image) *errors.RestErr
	CreateAlbum(album *models.Album) *errors.RestErr
	DeleteAlbum(album *models.Album) *errors.RestErr
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

func (ad albumDao) DeleteAlbum(album *models.Album) *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Collection("albums_list")
	filter := bson.D{{"_id", album.ID}}
	err := collection.FindOne(ctx, filter).Decode(&album)
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}
	//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = database.Collection(album.Email + "-" + album.AlbumName).Drop(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//Insert user ingto the database
func (ad albumDao) InsertImageInAlbum(image *models.Image) *errors.RestErr {
	bucket, err := gridfs.NewBucket(
		mongodb.GetMongoInstance().Database("Images"),
	)
	if err != nil {
		log.Fatal(err)
	}
	uploadStream, err := bucket.OpenUploadStream(
		image.ImageName, // this is the name of the file which will be saved in the database
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(image.ImageData)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(uploadStream.FileID)
	log.Printf("Write file to DB was successful. File size: %d \n", fileSize)

	var album models.Album
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Collection("albums_list")
	id, err := primitive.ObjectIDFromHex(image.AlbumId)
	filter := bson.D{{"_id", id}}
	if err := collection.FindOne(ctx, filter).Decode(&album); err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}
	fmt.Println("album")

	fmt.Println(uploadStream.FileID)
	image.FileId = uploadStream.FileID.(primitive.ObjectID).Hex()
	image.ImageData = nil
	fmt.Println("image")

	fmt.Println(image)
	collection = database.Collection(album.Email + "-" + album.AlbumName)
	_, err = collection.InsertOne(ctx, image)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

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
