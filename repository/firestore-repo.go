package repository

import (
	"context"
	"log"
	"task/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	projectId      string = "go-api-XXXXXX" // change your firebase project ID
	collectionName string = "posts"
)

type repo struct{}

// New Firestore Repository create a repository
func NewFirestoreRepository() TaskRepository {
	return &repo{}
}

func (r *repo) PostTask(task *model.Task) (*model.Task, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    task.ID,
		"Title": task.Title,
		"Text":  task.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add a new task: %v", err)
		return nil, err
	}

	return task, nil
}

func (r *repo) GetTaskList() ([]model.Task, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	var taskList []model.Task
	iter := client.Collection(collectionName).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of task: %v", err)
			return nil, err
		}
		task := model.Task{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		taskList = append(taskList, task)
	}

	return taskList, nil
}
