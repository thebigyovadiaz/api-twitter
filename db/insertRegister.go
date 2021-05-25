package db

import (
	"context"
	"time"

	"github.com/thebigyovadiaz/api-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertRegister is the func to save an user in DB */
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("api-twitter")
	col := db.Collection("users")

	u.Password, _ = PasswordEncrypt(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
