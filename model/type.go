package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Games struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string              `bson:"name,omitempty" json:"name,omitempty"`
	Rating       float64             `bson:"rating,omitempty" json:"rating,omitempty"`
	Release      primitive.DateTime  `bson:"release_date,omitempty" json:"release_date,omitempty"`
	Desc         string              `bson:"desc,omitempty" json:"desc,omitempty"`
	Genre        []string            `bson:"genre,omitempty" json:"genre,omitempty"`
	Dev_name     Developer           `bson:"dev_name,omitempty" json:"dev_name,omitempty"`
	Game_banner  string              `bson:"game_banner,omitempty" json:"game_banner,omitempty"`
	Preview      string              `bson:"preview,omitempty" json:"preview,omitempty"`
	Link_games	 string				 `bson:"link_games,omitempty" json:"link_games,omitempty"`
	Game_logo    string  			 `bson:"game_logo,omitempty" json:"game_logo,omitempty"`
}

type Developer struct {
	ID   primitive.ObjectID 		`bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             		`bson:"name,omitempty" json:"name,omitempty"`
	Bio  string             		`bson:"dev_bio,omitempty" json:"bio,omitempty"`
}

type Admin struct {
	ID        primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    User_name string             	`bson:"user_name,omitempty" json:"user_name,omitempty"`
	Name	  string				`bson:"name,omitempty" json:"name,omitempty"`
    Password  string             	`bson:"password,omitempty" json:"password,omitempty"`
	UpdatedAt time.Time          	`bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Token struct{
	ID			string 				`bson:"_id,omitempty" json:"_id,omitempty"`
	Token		string				`bson:"token" json:"token,omitempty"`
	AdminID		string				`bson:"admin_id" json:"admin_id,omitempty"`
	CreatedAt	time.Time			`bson:"created_at" json:"created_at"` 
}

type GoogleUser struct {
	ID            string `bson:"_id,omitempty" json:"_id,omitempty"`
	Email         string `bson:"email,omitempty" json:"email,omitempty"`
	VerifiedEmail bool   `bson:"verified_email" json:"verified_email"`
	Picture       string `bson:"picture" json:"picture"`
	Name          string `bson:"name" json:"name"`
}