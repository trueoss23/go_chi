package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/trueoss23/go_chi/models"
)

type MySQLDatabase struct {
	db *sql.DB
}

func (m *MySQLDatabase) Connect() error {
	dsn := "root:qwe@tcp(localhost:3306)/golang"
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

func (m *MySQLDatabase) Insert(data interface{}) error {
	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	book := data.(models.BookModel)
	_, err := m.db.Exec(query, book.Title, book.Author)
	if err != nil {
		return err
	}
	fmt.Println("Data inserted successfully")
	return nil
}

func (m *MySQLDatabase) Delete(id string) error {
	stmt, err := m.db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		return fmt.Errorf("не удалось подготовить запрос на удаление: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос на удаление: %w", err)
	}

	return nil
}

func (m *MySQLDatabase) Get(id string) (interface{}, error) {
	var book models.Book

	row := m.db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("книга с указанным ID не найдена")
	} else if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос на получение: %w", err)
	}

	return &book, nil
}

func (m *MySQLDatabase) GetAll() ([]interface{}, error) {
	var books []interface{}

	rows, err := m.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос на получение всех книг: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return nil, fmt.Errorf("не удалось прочитать результаты запроса: %w", err)
		}

		books = append(books, &book)
	}

	return books, nil
}
