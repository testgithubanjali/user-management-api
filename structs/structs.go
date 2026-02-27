package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}

type CreateUserPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateUserPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}