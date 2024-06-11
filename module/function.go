package module

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/rayfanaqbil/zenverse-BE/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

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
func InsertGames(name string, rating float64, desc string, genre []string, devname model.Developer, gamebanner string, preview string, gamelogo string)  (insertedID interface{}) {
	var datagame model.Games
	datagame.Name = name
	datagame.Rating = rating
	datagame.Release = primitive.NewDateTimeFromTime(time.Now().UTC())
	datagame.Desc = desc
	datagame.Genre = genre
	datagame.DevName = devname
	datagame.GameBanner = gamebanner
	datagame.Preview = preview
	datagame.GameLogo = gamelogo
	return InsertOneDoc("Zenverse", "Games", datagame)

}

func GetAllDataGames() (data []model.Games) {
	gem := MongoConnect("Zenverse").Collection("Games")
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