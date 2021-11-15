package posts

import (
	"article-service/models"
	"article-service/models/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostHandler) HandleCreate(c *gin.Context) {
	validated, _ := c.Get("validated_input")
	paramCreate := validated.(posts.ParameterCreatePost)

	errCreate := paramCreate.CreateNewPost()
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Error create : " + errCreate.Error(),
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
