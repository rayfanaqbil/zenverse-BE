package _zenverse

import (
	"fmt"
	"testing"

	"github.com/rayfanaqbil/zenverse-BE/model"
	"github.com/rayfanaqbil/zenverse-BE/module"
)

func TestInsertGames(t *testing.T) {
	name := "Resident Evil 2 Remake"
	rating := 10.0
	desc := "A deadly virus engulfs the residents of Raccoon City in September of 1998, plunging the city into chaos as flesh eating zombies roam the streets for survivors. An unparalleled adrenaline rush, gripping storyline, and unimaginable horrors await you. Witness the return of Resident Evil 2."
	genre := []string {"Action"}
	devname := model.Developer{
		Name: "CAPCOM Co., Ltd.",
		Bio:  "A remake of the original 1998 Resident Evil 2 game, it follows rookie police officer Leon S. Kennedy and student Claire Redfield as they attempt to escape from Raccoon City during a zombie apocalypse. It was released worldwide for the PlayStation 4, Windows, and Xbox One on January 25, 2019.",
	}
	gamebanner := "https://i.ibb.co.com/Jtbxq1Q/Resident-Evil-2-Remake-4-K-Wallpaper.jpg"
	preview := "https://www.youtube.com/watch?v=u3wS-Q2KBpk"
	gamelogo := "https://i.ibb.co.com/Rhwmct8/re2remake.png"

	insertedID, err := module.InsertGames(module.MongoConn, "Games", name, rating, desc, genre, devname, gamebanner, preview, gamelogo)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetAll(t *testing.T) {
	data := module.GetAllDataGames(module.MongoConn, "Games")
	fmt.Println(data)
}