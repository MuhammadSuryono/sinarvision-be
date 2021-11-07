package posts

import "article-service/db"

func (param ParameterUpdatePost) UpdatePost(postId int64) error {
	post := GetPostByID(postId)
	post.Title = param.Title
	post.Category = param.Category
	post.Status = param.Status
	post.Content = param.Content

	err := db.Connection.Save(&post)
	return err.Error
}
