package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	onceMongoDB sync.Once
	InstanceMongo *MongoDB
)

type MongoDB struct {
	Database *mongo.Database
}

func GetInstanceMongo() *MongoDB {

	onceMongoDB.Do(func() {
		InstanceMongo = &MongoDB{}

		//vars options
		user := "root"
		pass := "example"
		host := "localhost"
		port := 27017
		nameDB := "prueba"

		//Client Opts
		clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
		clientOpts.Auth = &options.Credential{
			Username:                user,
			Password:                pass,
		}

		//connect database mongo
		c, err := mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			log.Fatal(err)
		}


		InstanceMongo.Database = c.Database(nameDB)
	})


	// Check the connections
	err := InstanceMongo.Database.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return InstanceMongo
}

