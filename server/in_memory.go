package main

import (
	"fmt"
	pb "github.com/whipshout/grpc/proto/todo/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type inMemoryDB struct {
	tasks []*pb.Task
}

func (db *inMemoryDB) deleteTask(id uint64) error {
	for i, task := range db.tasks {
		if task.Id == id {
			db.tasks = append(db.tasks[:i], db.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

func (db *inMemoryDB) updateTask(id uint64, description string, dueDate time.Time, done bool) error {
	for i, task := range db.tasks {
		if task.Id == id {
			t := db.tasks[i]
			t.Description = description
			t.DueDate = timestamppb.New(dueDate)
			t.Done = done
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

func (db *inMemoryDB) getTasks(f func(interface{}) error) error {
	for _, task := range db.tasks {
		if err := f(task); err != nil {
			return err

		}
	}

	return nil
}

func (db *inMemoryDB) addTask(description string, dueDate time.Time) (uint64, error) {
	nextId := uint64(len(db.tasks) + 1)

	task := &pb.Task{
		Id:          nextId,
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}

	db.tasks = append(db.tasks, task)

	return nextId, nil
}

func New() Db {
	return &inMemoryDB{}
}
