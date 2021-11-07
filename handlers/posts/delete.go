package posts

import (
	"article-service/models"
	"article-service/models/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostHandler) HandleDelete(c *gin.Context) {
	var uri posts.UriPost
	_ = c.BindUri(&uri)

	errDelete := posts.DeletePost(uri.Id)
	if errDelete != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Error delete " + errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success delete post",
	})

}
