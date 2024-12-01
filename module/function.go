package module

import (
	"context"
	"log"
	"fmt"
	"errors"
	"time"
	"github.com/rayfanaqbil/zenverse-BE/v2/config"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertGames(db *mongo.Database, col string, name string, rating float64, desc string, genre []string, devname model.Developer, gamebanner string, preview string, linkgames string, gamelogo string) (insertedID primitive.ObjectID, err error) {
	games := bson.M{
	"name": name,
	"rating": rating,
	"release": primitive.NewDateTimeFromTime(time.Now().UTC()),
	"desc": desc,
	"genre": genre,
	"dev_name":  devname,
	"game_banner": gamebanner,
	"preview": preview,
	"link_games": linkgames,
	"game_logo": gamelogo,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), games)
	if err != nil {
		fmt.Printf("InsertGames: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllDataGames(db *mongo.Database, col string) (data []model.Games) {
	gem := db.Collection(col)
	filter := bson.M{}
	cursor, err := gem.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDataGames: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetGamesByName(db *mongo.Database, collection string, name string) ([]model.Games, error) {
    var games []model.Games
    filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}
    opts := options.Find().SetLimit(10)
    cursor, err := db.Collection(collection).Find(context.TODO(), filter, opts)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    
    if err := cursor.All(context.TODO(), &games); err != nil {
        return nil, err
    }
    
    return games, nil
}

func GetGamesByID(_id primitive.ObjectID, db *mongo.Database, col string) (games model.Games, errs error) {
	gem := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := gem.FindOne(context.TODO(), filter).Decode(&games)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return games, fmt.Errorf("no data found for ID %s", _id)
		}
		return games, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return games, nil
}

func UpdateGames(db *mongo.Database, col string, id primitive.ObjectID, name string, rating float64, desc string, genre []string, devname model.Developer, gamebanner string, preview string, linkgames string, gamelogo string) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":        name,
			"rating":      rating,
			"desc":        desc,
			"genre":       genre,
			"dev_name":    devname,
			"game_banner": gamebanner,
			"preview":     preview,
			"link_games":  linkgames,
			"game_logo":   gamelogo,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("UpdateGames: %v", err)
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}
	return nil
}

func DeleteGamesByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	gem := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := gem.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
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

func Login(db *mongo.Database, username string, password string) (string, error) {
    var admin model.Admin
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := db.Collection("Admin").FindOne(ctx, bson.M{"user_name": username}).Decode(&admin)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return "", fmt.Errorf("user not found")
        }
        return "", fmt.Errorf("error finding user: %v", err)
    }

    if admin.Password != password {
        return "", fmt.Errorf("invalid password")
    }

    token, err := config.GenerateJWT(admin)
    if err != nil {
        return "", fmt.Errorf("error generating token: %v", err)
    }

    return token, nil
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