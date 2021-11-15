package posts

import (
	"article-service/models"
	"article-service/models/posts"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func (cl *PostHandler) CreateAbleValidator(c *gin.Context) {
	var createParam posts.ParameterCreatePost
	opts := govalidator.Options{
		Data:            &createParam,
		Request:         c.Request,
		Rules:           createParam.GetValidationRules(),
		Messages:        createParam.GetValidationMessage(),
		RequiredDefault: true,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if e.Encode() == "" {
		c.Set("validated_input", createParam)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Error validate input",
			Data:      map[string]interface{}{"validation_error": e},
		})
	}
}

func (cl *PostHandler) UpdateAbleValidator(c *gin.Context) {
	var updateParam posts.ParameterUpdatePost
	opts := govalidator.Options{
		Data:            &updateParam,
		Request:         c.Request,
		Rules:           updateParam.GetValidationRules(),
		Messages:        updateParam.GetValidationMessage(),
		RequiredDefault: true,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if e.Encode() == "" {
		c.Set("validated_input", updateParam)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.CommonResponse{
			IsSuccess: false,
			Message:   "Error validate input",
			Data:      map[string]interface{}{"validation_error": e},
		})
	}
}
