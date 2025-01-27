package module

import(
"context"
	"log"
	"fmt"
	"errors"
	"time"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func GetGamesByHighestRating(db *mongo.Database, col string) ([]model.Games, error) {
    var games []model.Games
    filter := bson.M{} 
    opts := options.Find().SetSort(bson.M{"rating": -1})
    
    cursor, err := db.Collection(col).Find(context.TODO(), filter, opts)
    if err != nil {
        return nil, fmt.Errorf("error fetching games: %v", err)
    }
    defer cursor.Close(context.TODO())
    
    if err := cursor.All(context.TODO(), &games); err != nil {
        return nil, fmt.Errorf("error decoding games: %v", err)
    }
    
    return games, nil
}


