package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//User ...
type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName   string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	CreateDate string             `json:"create_date,omitempty" bson:"create_date,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
}
