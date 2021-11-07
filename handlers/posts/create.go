package posts

import (
	"article-service/models"
	"article-service/models/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostHandler) HandleCreate(c *gin.Context) {
	var paramCreate posts.ParameterCreatePost
	err := c.BindJSON(&paramCreate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Parameter request can't empty",
			Data:      paramCreate,
		})
		return
	}

	errCreate := paramCreate.CreateNewPost()
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Parameter request can't empty",
			Data:      paramCreate,
		})
		return
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		IsSuccess: true,
		Message:   "Success create post",
		Data:      paramCreate,
	})
	return

}
