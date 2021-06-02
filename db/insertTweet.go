package db

import (
	"context"
	"time"

	"github.com/thebigyovadiaz/api-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertTweet is the func to save a Tweet in DB */
func InsertTweet(t models.InsertTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConn.Database("api-twitter")
	col := db.Collection("tweets")

	register := bson.M{
		"userID":    t.UserID,
		"message":   t.Message,
		"createdAt": t.CreatedAt,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
