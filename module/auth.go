package module

import (
	"context"
	"fmt"
	"errors"
	"time"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"regexp"
)

func AddToBlacklist(db *mongo.Database, collection string, token string) error {
	blacklistEntry := bson.M{
		"token":     token,
		"createdAt": time.Now(),
	}
	_, err := db.Collection(collection).InsertOne(context.Background(), blacklistEntry)
	return err
}

func IsTokenBlacklisted(db *mongo.Database, collection string, token string) (bool, error) {
	var result bson.M
	err := db.Collection(collection).FindOne(context.Background(), bson.M{"token": token}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func validateUsername(username string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	if !re.MatchString(username) {
		return errors.New("invalid username format")
	}
	return nil
}

func InsertAdmin(db *mongo.Database, col string, username string, password string, email string) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
	"user_name" : username,
	"email"		: email,
	"password"	: password,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

	func DeleteTokenFromMongoDB(db *mongo.Database, col string, token string) error {
		collection := db.Collection(col)
		filter := bson.M{"token": token}

		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}

		return nil
	}

func SaveTokenToDatabase(db *mongo.Database, col string, adminID string, token string) error {
    collection := db.Collection(col)
    filter := bson.M{"admin_id": adminID}
    update := bson.M{
        "$set": bson.M{
            "token":      token,
            "updated_at": time.Now(),
        },
    }
    _, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
    if err != nil {
        return err
    }

    return nil
}

func GetAdminByUsername(db *mongo.Database, col string, username string) (*model.Admin, error) {
    if err := validateUsername(username); err != nil {
		return nil, err
	}
	var admin model.Admin
    err := db.Collection(col).FindOne(context.Background(), bson.M{"user_name": username}).Decode(&admin)
    if err == mongo.ErrNoDocuments {
        return nil, nil 
    }
    if err != nil {
        return nil, err
    }
    return &admin, nil
}

func GetAdminByEmail(db *mongo.Database, col string, email string) (*model.Admin, error) {
	var admin model.Admin
    err := db.Collection(col).FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)
    if err == mongo.ErrNoDocuments {
        return nil, nil 
    }
    if err != nil {
        return nil, err
    }
    return &admin, nil
}

func SaveGoogleUserToDatabase(db *mongo.Database, col string, googleUser model.GoogleUser) error {
    collection := db.Collection(col)

    filter := bson.M{"email": googleUser.Email}
    var existingUser model.GoogleUser
    err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
    if err == nil {
        return nil
    }
    if err != mongo.ErrNoDocuments {
        return err
    }

    _, err = collection.InsertOne(context.Background(), googleUser)
    if err != nil {
        return err
    }
    return nil
}