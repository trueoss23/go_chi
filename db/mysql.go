package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	cfg "github.com/trueoss23/go_chi/config"
	"github.com/trueoss23/go_chi/models"
)

type MySQLDatabase struct {
	db *sql.DB
}

func (m *MySQLDatabase) Connect() error {
	dsn := cfg.Cfg.DbUser + ":" + cfg.Cfg.DbPass + "@tcp(localhost:3306)/" + cfg.Cfg.DbName
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	m.db = conn
	return nil
}

func (m *MySQLDatabase) Close() error {
	return m.db.Close()
}

func (m *MySQLDatabase) Insert(book models.BookModel) error {
	m.Connect()
	defer m.db.Close()

	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	_, err := m.db.Exec(query, book.Title, book.Author)

	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLDatabase) Delete(id string) error {
	m.Connect()
	defer m.db.Close()

	elem, err := m.db.Exec("DELETE FROM books WHERE id = ?", id)
	fmt.Println(elem.RowsAffected())

	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос на удаление: %w", err)
	}

	return nil
}

func (m *MySQLDatabase) Get(id string) (models.Book, error) {
	var book models.Book
	m.Connect()
	defer m.db.Close()

	row := m.db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author)

	if err == sql.ErrNoRows {
		return models.Book{}, nil
	} else if err != nil {
		return models.Book{}, fmt.Errorf("не удалось выполнить запрос на получение: %w", err)
	}

	return book, nil
}

func (m *MySQLDatabase) GetAll() ([]models.Book, error) {
	var books []models.Book
	m.Connect()
	defer m.db.Close()

	rows, err := m.db.Query("SELECT id, title, author FROM books")

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
