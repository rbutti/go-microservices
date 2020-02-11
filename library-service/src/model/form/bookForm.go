package form

import (
	"library-service/model/domain"
	"time"
)

type BookForm struct {
	Title         string `json:"title" form:"required,max=255"`
	Author        string `json:"author" form:"required,alpha_space,max=255"`
	PublishedDate string `json:"published_date" form:"required,date"`
	ImageUrl      string `json:"image_url" form:"url"`
	Description   string `json:"description"`
}

func (f *BookForm) ToModel() (*domain.Book, error) {
	pubDate, err := time.Parse("2006-01-02", f.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &domain.Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageUrl:      f.ImageUrl,
		Description:   f.Description,
	}, nil
}
