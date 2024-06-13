package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Games struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string              `bson:"name,omitempty" json:"name,omitempty"`
	Rating       float64             `bson:"rating,omitempty" json:"rating,omitempty"`
	Release      primitive.DateTime  `bson:"release_date,omitempty" json:"release_date,omitempty"`
	Desc         string              `bson:"game_desc,omitempty" json:"game_desc,omitempty"`
	Genre        []string            `bson:"genre,omitempty" json:"genre,omitempty"`
	Dev_name     Developer           `bson:"dev_name,omitempty" json:"dev_name,omitempty"`
	Game_banner  string              `bson:"game_banner,omitempty" json:"game_banner,omitempty"`
	Preview      string              `bson:"preview,omitempty" json:"preview,omitempty"`
	Game_logo    string  			 `bson:"game_logo,omitempty" json:"game_logo,omitempty"`
}

type Developer struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Bio  string             `bson:"dev_bio,omitempty" json:"bio,omitempty"`
}
