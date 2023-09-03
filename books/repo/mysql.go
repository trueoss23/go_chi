package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/trueoss23/go_chi/domain/models"
)

type Repo interface {
	GetAll() ([]models.Book, error)
	Get(id string) (models.Book, error)
	Insert(book models.BookModel) (models.Book, error)
	Delete(id string) error
}

type MySQLRepo struct {
	ctx        context.Context
	connection *sql.DB
}

func NewMySQLRepo(ctx context.Context, connection *sql.DB) Repo {
	return &MySQLRepo{ctx: ctx, connection: connection}
}

func (m *MySQLRepo) Insert(book models.BookModel) (models.Book, error) {

	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	res, err := m.connection.Exec(query, book.Title, book.Author)

	if err != nil {
		_, cancel := context.WithCancel(m.ctx)
		cancel()
		return models.Book{}, err
	}
	id, _ := res.LastInsertId()
	resultBook := models.Book{
		ID:     strconv.FormatInt(id, 10),
		Title:  book.Title,
		Author: book.Author,
	}

	return resultBook, nil
}

func (m *MySQLRepo) Delete(id string) error {

	elem, err := m.connection.Exec("DELETE FROM books WHERE id = ?", id)
	fmt.Println(elem.RowsAffected())

	if err != nil {
		_, cancel := context.WithCancel(m.ctx)
		cancel()
		return fmt.Errorf("не удалось выполнить запрос на удаление: %w", err)
	}

	return nil
}

func (m *MySQLRepo) Get(id string) (models.Book, error) {
	var book models.Book

	row := m.connection.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author)

	if err == sql.ErrNoRows {
		_, cancel := context.WithCancel(m.ctx)
		cancel()
		return models.Book{}, nil
	} else if err != nil {
		return models.Book{}, fmt.Errorf("не удалось выполнить запрос на получение: %w", err)
	}

	return book, nil
}

func (m *MySQLRepo) GetAll() ([]models.Book, error) {
	var books []models.Book

	rows, err := m.connection.Query("SELECT id, title, author FROM books")

	if err != nil {
		_, cancel := context.WithCancel(m.ctx)
		cancel()
		return nil, fmt.Errorf("не удалось выполнить запрос на получение всех книг: %w", err)
	}

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)

		if err != nil {
			_, cancel := context.WithCancel(m.ctx)
			cancel()
			return nil, fmt.Errorf("не удалось прочитать результаты запроса: %w", err)
		}

		books = append(books, book)
	}
	return books, nil
}
