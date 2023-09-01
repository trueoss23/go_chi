package usecases

import (
	"errors"

	"github.com/trueoss23/go_chi/books/repo"
	
	"github.com/trueoss23/go_chi/models"
)

type Usecase interface {
	Connect() error
	Close() error
	GetAll() ([]models.Book, error)
	Get(id string) (models.Book, error)
	Insert(book models.BookModel) error
	Delete(id string) error
}

type BookUseCase struct {
	bookRepo repo.BookRepository
}

func NewBookUseCase(bookRepo repo.BookRepository) *BookUseCase {
	return &BookUseCase{
		bookRepo: bookRepo,
	}
}

func (uc *BookUseCase) CreateBook(title string, author string) (*models.Book, error) {
	if title == "" || author == "" {
		return nil, errors.New("Title and author cannot be empty")
	}

	book := &models.Book{
		Title:  title,
		Author: author,
	}

	err := uc.bookRepo.SaveBook(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *BookUseCase) GetBookByID(id string) (*models.Book, error) {
	book, err := uc.bookRepo.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *BookUseCase) UpdateBook(book *models.Book) error {
	if book == nil {
		return errors.New("Invalid book")
	}

	err := uc.bookRepo.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BookUseCase) DeleteBook(id string) error {
	err := uc.bookRepo.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}
