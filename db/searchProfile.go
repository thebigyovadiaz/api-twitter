package db

import (
	"context"
	"log"
	"time"

	"github.com/thebigyovadiaz/api-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* SearchProfile is the func to search a user in DB */
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConn.Database("api-twitter")
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		log.Println("User not found " + err.Error())
		return profile, err
	}

	return profile, nil
}
