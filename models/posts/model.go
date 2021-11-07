package posts

import "time"

type Post struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"size:200"`
	Content   string    `json:"content" gorm:"TEXT"`
	Category  string    `json:"category" gorm:"size:200"`
	Status    string    `json:"status" gorm:"size:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostList struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Content  string `json:"content" gorm:"TEXT"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type PostDetail struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content" gorm:"TEXT"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ParameterCreatePost struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" gorm:"TEXT" binding:"required"`
	Category string `json:"category" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

type ParameterUpdatePost struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" gorm:"TEXT" binding:"required"`
	Category string `json:"category" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

type UriPost struct {
	Id int64 `uri:"id"`
}

type QueryRequest struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type ParameterSearchPost struct {
	Title    string `form:"title"`
	Category string `form:"category"`
	Status   string `form:"status"`
}
