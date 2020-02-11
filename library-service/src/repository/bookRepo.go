package repository

import (
	"github.com/jinzhu/gorm"
	"library-service/model/domain"
)

func ListBooks(db *gorm.DB) (domain.Books, error) {
	books := make([]*domain.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, nil
	}

	return books, nil
}

func CreateBook(db *gorm.DB, book *domain.Book) (*domain.Book, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func ReadBook(db *gorm.DB, id uint) (*domain.Book, error) {
	book := &domain.Book{}
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(db *gorm.DB, book *domain.Book) error {
	if err := db.First(&domain.Book{}, book.ID).Update(book).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBook(db *gorm.DB, id uint) error {
	book := &domain.Book{}
	if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
