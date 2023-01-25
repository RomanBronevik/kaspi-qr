package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
}

func NewRepository() *Repository {
	return &Repository{}
}
