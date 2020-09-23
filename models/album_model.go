package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Info ...
type Info struct {
	CreatedBy   string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
}

//Album ...
type Album struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AlbumName string             `json:"album_name,omitempty" bson:"album_name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Info      Info               `json:"info,omitempty" bson:"info,omitempty"`
}

//Image ...
type Image struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AlbumId   string             `json:"album_id,omitempty" bson:"album_id,omitempty"`
	ImageName string             `json:"image_name,omitempty" bson:"image_name,omitempty"`
	ImageData []byte             `json:"image_data,omitempty" bson:"image_data,omitempty"`
	FileId    string             `json:"file_id,omitempty" bson:"file_id,omitempty"`
	Info      Info               `json:"info,omitempty" bson:"info,omitempty"`
}
