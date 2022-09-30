package conversion

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetMongoId(id string) primitive.ObjectID {
	iddd, _ := primitive.ObjectIDFromHex(id)
	return iddd
}

func GetIdFromMongoId(mongoId interface{}) string {
	stringObjectID := mongoId.(primitive.ObjectID).Hex()
	return stringObjectID
}
