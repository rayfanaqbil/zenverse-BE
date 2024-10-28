package module

import(
	"context"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func IsAdmin(db *mongo.Database, col string, email string) bool {
	var admin model.Admin
	err := db.Collection(col).FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)
	return err == nil
}