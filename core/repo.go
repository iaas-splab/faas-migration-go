package core

type Repository interface {
	Init() error
	Put(*ToDoItem) (*ToDoItem, error)
	Get(string) (*ToDoItem, error)
	List() ([]ToDoItem, error)
	Delete(string) error
}
