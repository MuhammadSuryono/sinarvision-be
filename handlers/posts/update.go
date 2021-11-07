package posts

import (
	"article-service/models"
	"article-service/models/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostHandler) HandleUpdate(c *gin.Context) {
	var paramUpdate posts.ParameterUpdatePost
	var uri posts.UriPost
	_ = c.BindUri(&uri)
	err := c.BindJSON(&paramUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Parameter request can't empty",
			Data:      paramUpdate,
		})
		return
	}

	errCreate := paramUpdate.UpdatePost(uri.Id)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Error update: " + errCreate.Error(),
			Data:      paramUpdate,
		})
		return
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success update post",
		Data:      paramUpdate,
	})
	return

}
