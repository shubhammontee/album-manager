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
	AlbumName string             `json:"album_name,omitempty"`
	Email     string             `json:"email,omitempty"`
	Info      Info               `json:"info,omitempty"`
}

//Image ...
type Image struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ImageName string             `json:"image_name,omitempty"`
	Info      Info               `json:"info,omitempty"`
}
