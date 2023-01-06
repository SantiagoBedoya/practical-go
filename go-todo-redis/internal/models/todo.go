package models

// Todo define data struct for todo
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

type (
	// TodoService define the behavior for todo services
	TodoService interface {
		FindAll() ([]Todo, error)
		FindByID(todoID string) (*Todo, error)
		Create(todo Todo) error
	}
)
