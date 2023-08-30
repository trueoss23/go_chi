package db

type Database interface {
    Connect() error
    Close() error
    Insert(data interface{}) error
    Delete(id string) error
    Get(id string) (interface{}, error)
    GetAll() ([]interface{}, error)
}