package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTask
// Creates a new task in the database.
func (d *database) CreateTask(task models.Task) (*models.Task, error) {
	var (
		result models.Task
	)
	task.Id = uuid.NewString()
	task.Status = models.TaskStatusPending
	_, err := d.TaskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}
	err = d.TaskCollection.FindOne(context.TODO(), bson.D{{"_id", task.Id}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, utils.ErrTaskNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTask
// Returns a task from the database based on the specified id.
func (d *database) GetTask(id string) (*models.Task, error) {
	var result models.Task
	err := d.TaskCollection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, utils.ErrTaskNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListTasks
// Returns tasks from the database based on the supplied filters.
func (d *database) ListTasks(action string, character string, limit int64, status string) (*[]models.Task, error) {
	var tasks []models.Task
	opts := options.Find().SetLimit(limit)
	filter := bson.D{}
	if action != "" {
		filter = append(filter, bson.E{"action", action})
	}
	if character != "" {
		filter = append(filter, bson.E{"character", character})
	}
	if status != "" {
		filter = append(filter, bson.E{"status", status})
	}
	cursor, err := d.TaskCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(tasks) == 0 {
		return nil, utils.ErrTasksNotFound
	}
	return &tasks, nil
}

// UpdateTask
// Updates an existing task in the database.
func (d *database) UpdateTask(id string, reason string, status string) error {
	filter := bson.D{
		{"_id", id},
	}
	update := bson.D{
		{"$set", bson.D{
			{"status", status},
			{"errorReason", reason},
		}},
	}
	err := d.TaskCollection.FindOneAndUpdate(context.TODO(), filter, update).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return utils.ErrTaskNotFound
	} else if err != nil {
		return err
	}
	return nil
}
