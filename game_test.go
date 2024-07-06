package _zenverse

import (
	"fmt"
	"testing"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
	"github.com/rayfanaqbil/zenverse-BE/v2/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertGames(t *testing.T) {
	name := "Genshin Impact"
	rating := 9.10
	desc := "Step into Teyvat, a vast world teeming with life and flowing with elemental energy. You and your sibling arrived here from another world. Separated by an unknown god, stripped of your powers, and cast into a deep slumber, you now awake to a world very different from when you first arrived. Thus begins your journey across Teyvat to seek answers from The Seven — the gods of each element. Along the way, prepare to explore every inch of this wondrous world, join forces with a diverse range of characters, and unravel the countless mysteries that Teyvat holds..."
	genre := []string{"Adventure game", "Action role-playing game"}
	devname := model.Developer{
		Name: "HoYoverse",
		Bio:  "tech otaku save the world",
	}
	gamebanner := "https://i.ibb.co.com/k1KdV7t/genshin-main-banner.png"
	preview := "https://www.youtube.com/watch?v=qqnEjmnitgc"
	linkgames := "https://genshin.hoyoverse.com/id/" 
	gamelogo := "https://i.ibb.co.com/Z6xFZP6/genshin-logo.png"
	insertedID, err := module.InsertGames(module.MongoConn, "Games", name, rating, desc, genre, devname, gamebanner, preview, linkgames, gamelogo)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetAll(t *testing.T) {
	data := module.GetAllDataGames(module.MongoConn, "Games")
	fmt.Println(data)
}

func TestGetGamesByID(t *testing.T) {
	id := "666b19daa1296db477837ee9"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	geming, err := module.GetGamesByID(objectID, module.MongoConn, "Games")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(geming)
}

func TestDeletePresensiByID(t *testing.T) {
	id := "6412ce78686d9e9ba557cf8a" 
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteGamesByID(objectID, module.MongoConn, "Games")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	_, err = module.GetGamesByID(objectID, module.MongoConn, "Games")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestInsertAdmin(t *testing.T) {
	username 	 := "Zenverse"
	password 	 := "zenverse123"
	insertedID, err := module.InsertAdmin(module.MongoConn, "Admin", username, password)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestLogin(t *testing.T) {
	username := "Zenverse"
	password := "zenverse123"
	token, err := module.Login(module.MongoConn, username, password)
	if err != nil {
		t.Errorf("Login failed: %v", err)
	} else {
		fmt.Printf("Login successful. Token: %s\n", token)
	}
}

func TestGetGamesByName(t *testing.T) {
	name := "Resident"
	games, err := module.GetGamesByName(module.MongoConn, "Games", name)
	if err != nil {
		t.Fatalf("Error calling GetGamesByName: %v", err)
	}

	fmt.Println("Game found:", games)
}

func TestGetAdminByUs(t *testing.T) {
	username := "Zenverse"
	admins, err := module.GetAdminByUsername(module.MongoConn, "Admin", username)
	if err != nil {
		t.Fatalf("Error calling GetByAdminUsername: %v", err)
	}

	fmt.Println("Username Found:", admins)
}
