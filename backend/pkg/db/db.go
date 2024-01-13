package db

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo/pkg/model"
)

type DB struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	connectionURI  = "mongodb://localhost:27017"
	dbName         = "todoapp"
	collectionName = "todoapp"
)

func NewDB() *DB {
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	d := client.Database(dbName)
	collection := d.Collection(collectionName)

	log.Infof("DB Connected Successfully")
	return &DB{
		client:     client,
		db:         d,
		collection: collection,
	}
}

func (d *DB) CreateTask(t *model.Task) (string, error) {
	res, err := d.collection.InsertOne(context.TODO(), t)
	if err != nil {
		return "", err
	}
	id, _ := res.InsertedID.(string)
	return id, nil
}

func (d *DB) DeleteTask(id string) (int64, error) {
	filter := bson.M{"_id": id}
	res, err := d.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (d *DB) GetTask(id string) (*model.Task, error) {
	filter := bson.M{"_id": id}
	task := &model.Task{}
	err := d.collection.FindOne(context.TODO(), filter).Decode(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (d *DB) MarkTaskDone(id string, done bool) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"done": done}}

	_, err := d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetAllTask(owner string) ([]*model.Task, error) {
	filter := bson.M{"owner": owner}
	var tasks []*model.Task

	cursor, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		task := &model.Task{}
		err := cursor.Decode(task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
