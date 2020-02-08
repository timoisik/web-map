package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net"
	"time"
)

type Domain struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	Tld string
	Ips []net.IP
	CreatedAt time.Time
	FetchedAt time.Time
}

func (d *Domain) HasBeenChecked() bool {
	return d.FetchedAt.IsZero()
}

func (d *Domain) GetUrl() string {
	return d.Name + "." + d.Tld
}

func (d *Domain) Create() *mongo.InsertOneResult {
	model, err := Db.InsertOne(context.TODO(), d)

	if err != nil {
		log.Fatal(err)
	}

	d.Id = model.InsertedID.(primitive.ObjectID)
	return model
}

func (d *Domain) Read(b bson.D) Domain {
	var result Domain
	err := Db.FindOne(context.TODO(), b).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (d *Domain) Update(b bson.D) {
	filter := bson.D{{"_id", d.Id}}
	Db.UpdateOne(context.TODO(), filter, b)
}

func (d *Domain) Delete() {
	filter := bson.D{{"_id", d.Id}}
	Db.DeleteOne(context.TODO(), filter)
}