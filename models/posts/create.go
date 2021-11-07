package posts

import (
	"article-service/db"
)

func (param ParameterCreatePost) CreateNewPost() error {
	newData := Post{
		Title:    param.Title,
		Content:  param.Content,
		Category: param.Category,
		Status:   param.Status,
	}

	err := db.Connection.Create(&newData)
	return err.Error
}
