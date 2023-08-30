package db

type Database interface {
	Connect() error
	Close() error
	GetAll() ([]interface{}, error)
	Get(id string) (interface{}, error)
	Insert(data interface{}) error
	Delete(id string) error
}
