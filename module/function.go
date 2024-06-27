package module

import (
	"context"
	"fmt"
	"errors"
	"time"
	"github.com/rayfanaqbil/zenverse-BE/model"
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

func GetGamesByName(db *mongo.Database, col string, name string,) (games model.Games, errs error) {
	gem := db.Collection(col)
	filter := bson.M{"name": name}
	err := gem.FindOne(context.TODO(), filter).Decode(&games)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return games, fmt.Errorf("no data found for Name %s", name)
		}
		return games, fmt.Errorf("error retrieving data for Name %s: %s", name, err.Error())
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

func UpdateGames(db *mongo.Database, col string, id primitive.ObjectID, name string, rating float64, desc string, genre []string, devname model.Developer, gamebanner string, preview string, linkgames string, gamelogo string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":    name,
			"rating":     rating,
			"desc":     desc,
			"genre": genre,
			"dev_name": devname,
			"game_banner": gamebanner,
			"preview":  preview,
			"link_games": linkgames,
			"game_logo": gamelogo,
		},
}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateGames: %v\n", err)
		return err
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
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

func InsertAdmin(db *mongo.Database, col string, username string, password string ) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
	"user_name" : username,
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

func Login(db *mongo.Database, col string, username string, password string) (model.Admin, error) {
    var User model.Admin
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := db.Collection(col).FindOne(ctx, bson.M{"user_name": username}).Decode(&User)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return model.Admin{}, fmt.Errorf("user not found")
        }
        return model.Admin{}, fmt.Errorf("error finding user: %v", err)
    }

    if User.Password != password {
        return model.Admin{}, fmt.Errorf("invalid password")
    }

    return User, nil
}

func GetDataAdmin(db *mongo.Database, col string) (data []model.Admin) {
	admin := db.Collection(col)
	filter := bson.M{}
	cursor, err := admin.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAdmin: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
