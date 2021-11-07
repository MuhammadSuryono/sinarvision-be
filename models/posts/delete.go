package posts

import "article-service/db"

func DeletePost(postId int64) error {
	post := GetPostByID(postId)
	post.Status = "Thrash"

	err := db.Connection.Save(&post)
	return err.Error
}
