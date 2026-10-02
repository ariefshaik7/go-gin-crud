package users

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string        `json:"name" bson:"name"`
	Email     string        `json:"email" bson:"email"`
	Phone     string        `json:"phone" bson:"phone"`
	Age       int           `bson:"age" json:"age"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}
