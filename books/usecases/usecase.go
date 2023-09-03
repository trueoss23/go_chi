package usecases

import (
	"github.com/trueoss23/go_chi/books/repo"

	"github.com/trueoss23/go_chi/domain/models"
)

type Usecase interface {
	GetAll() ([]models.Book, error)
	Get(id string) (models.Book, error)
	Insert(book models.BookModel) (models.Book, error)
	Delete(id string) error
}

type BookUseCase struct {
	bookRepo repo.Repo
}

func NewBookUseCase(bookRepo repo.Repo) Usecase {
	return &BookUseCase{
		bookRepo: bookRepo,
	}
}

func (uc *BookUseCase) GetAll() ([]models.Book, error) {
	books, err := uc.bookRepo.GetAll()
	if err != nil {
		return []models.Book{}, err
	}
	return books, nil
}

func (uc *BookUseCase) Insert(bookmodel models.BookModel) (models.Book, error) {
	bookInsert, err := uc.bookRepo.Insert(bookmodel)
	if err != nil {
		return models.Book{}, err
	}
	return bookInsert, nil
}

func (uc *BookUseCase) Get(id string) (models.Book, error) {
	book, err := uc.bookRepo.Get(id)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (uc *BookUseCase) Delete(id string) error {
	err := uc.bookRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
