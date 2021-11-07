package posts

import "github.com/gin-gonic/gin"

type IPost interface {
	HandleCreate(c *gin.Context)
	HandleDelete(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleListPost(c *gin.Context)
	HandleListPostPublish(c *gin.Context)
	HandleReadPost(c *gin.Context)
	HandleCountStatus(c *gin.Context)
}

type PostHandler struct {
}

func NewPostHandler() IPost {
	return &PostHandler{}
}
