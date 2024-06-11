package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Games struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string              `bson:"name,omitempty" json:"name,omitempty"`
	Rating      float64             `bson:"rating,omitempty" json:"rating,omitempty"`
	Release     primitive.DateTime  `bson:"release_date,omitempty" json:"release_date,omitempty"`
	Desc        string              `bson:"game_desc,omitempty" json:"game_desc,omitempty"`
	Genre       []string            `bson:"genre,omitempty" json:"genre,omitempty"`
	DevName     Developer           `bson:"developer,omitempty" json:"developer,omitempty"`
	GameBanner  string              `bson:"game_banner,omitempty" json:"game_banner,omitempty"`
	Preview     string              `bson:"preview,omitempty" json:"preview,omitempty"`
	GameLogo    string  			`bson:"game_logo,omitempty" json:"game_logo,omitempty"`
}

type Developer struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Bio  string             `bson:"dev_bio,omitempty" json:"dev_bio,omitempty"`
}
