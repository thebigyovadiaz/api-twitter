package db

import (
	"context"
	"time"

	"github.com/thebigyovadiaz/api-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModifiedRegister is a func to modified user register */
func ModifiedRegister(user *models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConn.Database("api-twitter")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(user.Firstname) > 0 {
		register["firstname"] = user.Firstname
	}

	if len(user.Lastname) > 0 {
		register["lastname"] = user.Lastname
	}

	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	if len(user.Location) > 0 {
		register["location"] = user.Location
	}

	if len(user.Website) > 0 {
		register["website"] = user.Website
	}

	register["birthday"] = user.Birthday

	conditionUpdate := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filterUpdated := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := col.UpdateOne(ctx, filterUpdated, conditionUpdate)
	if err != nil {
		return false, err
	}

	return true, nil
}
