package posts

import (
	"article-service/db"
	"article-service/pagination"
	"gorm.io/gorm"
)

func GetPostByID(postId int64) (post Post) {
	db.Connection.Where("id = ?", postId).First(&post)
	return
}

func (uriArticle UriPaging) AllPostList(paramSearch ParameterSearchPost) *pagination.Paginator {
	query := paramSearch.queryAllPost()
	paginator := pagination.Paging(&pagination.Param{
		DB:      query,
		Limit:   int(uriArticle.Limit),
		Offset:  int(uriArticle.Offset),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &[]Post{})

	return paginator
}

func (p ParameterSearchPost) queryAllPost() *gorm.DB {
	query := db.Connection
	if p.Title != "" {
		query = query.Where("title LIKE ?", "%"+p.Title+"%")
	}

	if p.Category != "" {
		query = query.Where("category = ?", p.Category)
	}

	if p.Status != "" {
		query = query.Where("status = ?", p.Status)
	}

	return query
}

func (param ParameterSearchPost) CountDataByStatus() (total int64) {
	db.Connection.Model(&Post{}).Where("status = ?", param.Status).Count(&total)
	return
}
