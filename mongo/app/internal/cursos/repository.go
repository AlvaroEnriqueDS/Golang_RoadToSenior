package cursos

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"mongo/app/internal/core/domain"
	"mongo/app/internal/mongo"
)

type cursoRepository struct {}

func NewCursoRepository() *cursoRepository {
	return  new(cursoRepository)
}

func (r *cursoRepository) GetRepository(id string) (domain.Cursos,error){
	var person domain.Cursos

	//person = r.dataMemory[id]
	return person, nil
}

func (r *cursoRepository) SaveRepository(curso *domain.Cursos) error{
	c := mongo.GetInstanceMongo().Database.Collection("Cursos")
	insertResult, err := c.InsertOne(context.TODO(), curso)
	if err != nil {
		//log.Fatal(err)
		return err
	}

	fmt.Println("created: ", insertResult.InsertedID)

	return nil
}

func (r *cursoRepository) GetByNameRepository(name string) ([]domain.Cursos, error) {
	c := GetInstanceMongo().Database.Collection("Cursos")

	filter := bson.D{{"nombre", name}}
	//filter := bson.M{"$text": bson.M{"$search": name}}
	//{'$text': {'$search': nombre}}
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var data []domain.Cursos
	for cursor.Next(context.TODO()) {
		p := domain.Cursos{}
		err := cursor.Decode(&p)
		if err != nil {
			return nil, err
		}

		data = append(data, p)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	fmt.Println(data)

	return data, nil
}
