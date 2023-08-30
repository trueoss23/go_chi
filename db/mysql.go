package db

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct {
    db *sql.DB
}

func NewMySQLDatabase(connectionString string) (*MySQLDatabase, error) {
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
    }

    return &MySQLDatabase{
        db: db,
    }, nil
}

func (m *MySQLDatabase) Connect() error {
    // Уже подключено в NewMySQLDatabase()
    return nil
}

func (m *MySQLDatabase) Close() error {
    return m.db.Close()
}

func (m *MySQLDatabase) Insert(data interface{}) error {
    book, ok := data.(*Book)
    if !ok {
        return fmt.Errorf("неправильный тип данных")
    }

    stmt, err := m.db.Prepare("INSERT INTO books (id, title, author) VALUES (?, ?, ?)")
    if err != nil {
        return fmt.Errorf("не удалось подготовить запрос на вставку: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(book.ID, book.Title, book.Author)
    if err != nil {
        return fmt.Errorf("не удалось выполнить запрос на вставку: %w", err)
    }

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
    var book Book

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
        var book Book
        err := rows.Scan(&book.ID, &book.Title, &book.Author)
        if err != nil {
            return nil, fmt.Errorf("не удалось прочитать результаты запроса: %w", err)
        }

        books = append(books, &book)
    }

    return books, nil
}
