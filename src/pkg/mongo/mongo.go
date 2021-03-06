package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"app/src/pkg/conf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Setup() (err error) {
	var uri string
	if conf.Mongodb.User != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Mongodb.User, conf.Mongodb.Password, conf.Mongodb.Host, conf.Mongodb.Port)
	} else {
		uri = fmt.Sprintf("mongodb://@%s:%d", conf.Mongodb.Host, conf.Mongodb.Port)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("new client err:%v", err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	DB = client.Database(conf.Mongodb.Name)
	if err != nil {
		fmt.Printf("client connect err:%v", err)
		return err
	}
	return err
}

func InsertOne(name string, data interface{}) (id interface{}, err error) {
	res, err := DB.Collection(name).InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func One(name string, filter interface{}) (model interface{}, err error) {
	err = DB.Collection(name).FindOne(context.Background(), filter).Decode(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func All(name string, filter interface{}) (list []interface{}, err error) {
	cur, err := DB.Collection(name).Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	count, _ := DB.Collection(name).CountDocuments(context.Background(), filter)
	list = make([]interface{}, count)
	for cur.Next(context.Background()) {
		list = append(list, cur.Current)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return list, err
}
