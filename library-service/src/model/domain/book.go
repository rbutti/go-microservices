package domain

import (
	"github.com/jinzhu/gorm"
	"library-service/model/dto"
	"time"
)

type Books []*Book

type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	ImageUrl      string
	Description   string
}

func (b Book) ToDto() *dto.BookDto {
	return &dto.BookDto{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageUrl:      b.ImageUrl,
		Description:   b.Description,
	}
}

func (bs Books) ToDto() dto.BookDtos {
	dtos := make([]*dto.BookDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}
