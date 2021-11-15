package posts

import (
	"article-service/models"
	"article-service/models/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostHandler) HandleListPost(c *gin.Context) {
	var querySearch posts.ParameterSearchPost
	var uriArticle posts.UriPaging
	_ = c.BindQuery(&querySearch)
	_ = c.BindUri(&uriArticle)

	data := uriArticle.AllPostList(querySearch)
	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success retrieve data",
		Data:      data,
	})
}

func (p *PostHandler) HandleListPostPublish(c *gin.Context) {
	var querySearch posts.ParameterSearchPost
	var uriArticle posts.UriPaging
	_ = c.BindQuery(&querySearch)
	_ = c.BindUri(&uriArticle)

	querySearch.Status = "publish"
	data := uriArticle.AllPostList(querySearch)
	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success retrieve data",
		Data:      data,
	})
}

func (p *PostHandler) HandleReadPost(c *gin.Context) {
	var uri posts.UriPost
	_ = c.BindUri(&uri)

	data := posts.GetPostByID(uri.Id)
	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success retrieve data",
		Data:      data,
	})
}

func (p *PostHandler) HandleCountStatus(c *gin.Context) {
	var querySearch posts.ParameterSearchPost
	_ = c.BindQuery(&querySearch)

	total := querySearch.CountDataByStatus()
	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success retrieve data",
		Data:      total,
	})
}
