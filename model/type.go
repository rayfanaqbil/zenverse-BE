package model

import (
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
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Bio  string             `bson:"dev_bio,omitempty" json:"bio,omitempty"`
}

type Credentials struct {
    Username string 				`bson:"user_name,omitempty" json:"username,omitempty"`
    Password string 				`bson:"password,omitempty" json:"password,omitempty"`
}

type Admin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    User_name string             `bson:"user_name,omitempty" json:"user_name,omitempty"`
    Password  string             `bson:"password,omitempty" json:"password,omitempty"`
    Token     string             `bson:"token,omitempty" json:"token,omitempty"`
}
