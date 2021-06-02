package models

/* Tweet is the structure's Tweet */
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
