package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Warframe struct {
	Id             primitive.ObjectID   `json:"id" bson:"_id"`
	Name           string               `json:"name"`
	ImageName      string               `json:"image_name"`
	Description    string               `json:"description"`
	Health         int                  `json:"health"`
	Shield         int                  `json:"shield"`
	Armor          int                  `json:"armor"`
	Power          int                  `json:"power"`
	Sprint         float64              `json:"sprint"`
	MasteryReq     int                  `json:"mastery_req"`
	BuildPrice     int                  `json:"build_price"`
	BuildTime      int                  `json:"build_time"`
	BuildSkipPrice int                  `json:"build_skip_price" bson:"skipBuildTimePrice"`
	Components     []*WarframeComponent `json:"components"`
	Abilities      []*WarframeAbility   `json:"abilities"`
}

type WarframeComponent struct {
	Name      string          `json:"name" bson:"name"`
	ItemCount int             `json:"item_count"`
	Drops     []*WarframeDrop `json:"drops"`
}

type WarframeDrop struct {
	Location string  `json:"location"`
	Chance   float64 `json:"chance"`
	Rarity   string  `json:"rarity"`
}

type WarframeAbility struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FindWarframeNormals() []*Warframe {
	var result []*Warframe

	cur, err := WarframeNormalCol.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err.Error())
	}

	if err = cur.All(context.TODO(), &result); err != nil {
		log.Println(err.Error())
	}

	return result
}

func FindWarframePrimes() []*Warframe {
	var result []*Warframe

	cur, err := WarframePrimeCol.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err.Error())
	}

	if err = cur.All(context.TODO(), &result); err != nil {
		log.Println(err.Error())
	}

	return result
}
