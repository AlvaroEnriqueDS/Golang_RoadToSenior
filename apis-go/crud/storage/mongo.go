package storage

import (
	model "apis-go/crud/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo .
type Mongo struct {
	currentID int
	Persons   map[int]model.Person
	Database *mongo.Database
}

// NewMemory .
func NewMongo() Mongo {

	//vars options
	user := "root"
	pass := "example"
	host := "localhost"
	port := 27017

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

	// Check the connections
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}



	return Mongo{
		currentID: 0,
		Persons:   nil,
		Database:   c.Database("prueba"),
	}

	//persons := make(map[int]model.Person)
}

// Create .
func (m *Mongo) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	c := m.Database.Collection("Person")
	insertResult, err := c.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("created: ", insertResult.InsertedID)

	return nil
}

// Update actualiza una persona en el slice de memoria
func (m *Mongo) Update(ID string, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	c := m.Database.Collection("Person")


	objectD, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	//vacia id person
	person.ID = ""

	filter := bson.M{"_id": bson.M{"$eq": objectD}}
	//filter := bson.M{"_id": objectD}
	//upt := bson.D{{"$set", bson.D{{"age", person.Age}}}}
	upt := bson.D{{"$set", person}}
	updateResult, err := c.UpdateOne(context.TODO(), filter, upt)
	if err != nil {
		log.Fatal(err)
		return model.ErrIDPersonDoesNotExists
	}

	fmt.Println("update: ", updateResult.UpsertedID)

	return nil
}

// Delete borra de la memoria la persona
func (m *Mongo) Delete(ID string) error {

	c := m.Database.Collection("Person")

	objectD, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": objectD}}
	deleteResult, err := c.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return model.ErrIDPersonDoesNotExists
	}

	fmt.Println("delete: ", deleteResult.DeletedCount)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Mongo) GetByID(ID int) (model.Person, error) {
	c := m.Database.Collection("Person")
	filter := bson.D{{"_id", ID}}
	singleResult := c.FindOne(context.TODO(), filter)

	p := model.Person{}
	err := singleResult.Decode(&p)
	if err != nil {
		return p, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return p, nil
}

// GetAll retorna todas las personas que están en la memoria
func (m *Mongo) GetAll() (model.Persons, error) {
	var result model.Persons

	c := m.Database.Collection("Person")
	filter := bson.D{{}}
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		return result, fmt.Errorf("ID: %d: %w", 1, model.ErrIDPersonDoesNotExists)
	}

	for cursor.Next(context.TODO()) {
		p := model.Person{}
		err := cursor.Decode(&p)
		if err != nil {
			return result, fmt.Errorf("ID: %d: %w", 2, model.ErrIDPersonDoesNotExists)
		}

		result = append(result, p)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return result, fmt.Errorf("ID: %d: %w", 3, model.ErrIDPersonDoesNotExists)
	}


	return result, nil
}