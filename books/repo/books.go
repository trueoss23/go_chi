package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"ithub.com/trueoss23/go_chi/domain/models"
)

type Repo interface {
	GetAll() ([]models.Book, error)
	Get(id string) (models.Book, error)
	Insert(book models.BookModel) error
	Delete(id string) error
}

type MySQLRepo struct {
	connection *sql.DB
}

func NewMySQLRepo(connection *sql.DB) *MySQLRepo {
	return MySQLRepo{connection: connection}
}

func (m *MySQLRepo) Insert(book models.BookModel) error {
	// defer m.connection.Close()

	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	_, err := m.connection.Exec(query, book.Title, book.Author)

	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLRepo) Delete(id string) error {
	// m.Connect()
	// defer m.connection.Close()

	elem, err := m.connection.Exec("DELETE FROM books WHERE id = ?", id)
	fmt.Println(elem.RowsAffected())

	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос на удаление: %w", err)
	}

	return nil
}

func (m *MySQLRepo) Get(id string) (models.Book, error) {
	var book models.Book
	// m.Connect()
	// defer m.connection.Close()

	row := m.connection.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author)

	if err == sql.ErrNoRows {
		return models.Book{}, nil
	} else if err != nil {
		return models.Book{}, fmt.Errorf("не удалось выполнить запрос на получение: %w", err)
	}

	return book, nil
}

func (m *MySQLRepo) GetAll() ([]models.Book, error) {
	var books []models.Book
	// m.Connect()
	// defer m.connection.Close()

	rows, err := m.connection.Query("SELECT id, title, author FROM books")

	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос на получение всех книг: %w", err)
	}

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)

		if err != nil {
			return nil, fmt.Errorf("не удалось прочитать результаты запроса: %w", err)
		}

		books = append(books, book)
	}
	return books, nil
}
