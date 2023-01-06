package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SantiagoBedoya/todo-app/internal/models"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

type todoService struct {
	rdb *redis.Client
}

func NewTodoService(rdb *redis.Client) models.TodoService {
	return &todoService{rdb}
}

func (s todoService) FindAll() ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	todos := make([]models.Todo, 0)
	result, err := s.rdb.Get(ctx, "todos").Result()
	if err != nil {
		if err == redis.Nil {
			return todos, nil
		}
		return nil, err
	}
	if err := json.Unmarshal([]byte(result), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (s todoService) Create(todo models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	currentTodos, err := s.FindAll()
	if err != nil {
		return err
	}
	todo.ID = uuid.NewString()
	todo.IsCompleted = false
	currentTodos = append(currentTodos, todo)
	bytes, _ := json.Marshal(currentTodos)
	if err := s.rdb.Set(ctx, "todos", string(bytes), 10*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (s todoService) FindByID(todoID string) (*models.Todo, error) {
	currentTodos, err := s.FindAll()
	if err != nil {
		return nil, err
	}
	for _, todo := range currentTodos {
		if todo.ID == todoID {
			return &todo, nil
		}
	}
	return nil, nil
}
