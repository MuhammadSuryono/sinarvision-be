package main

import (
	"article-service/db"
	"article-service/handlers/posts"
	"article-service/models"
	postModel "article-service/models/posts"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	db.Connection.AutoMigrate(&postModel.Post{})

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, models.CommonResponse{
			IsSuccess: false,
			Message:   "Method Not Allowed",
		})
		return
	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, models.CommonResponse{
			IsSuccess: false,
			Message:   "Route Not Found",
		})
		return
	})

	articleHandler := posts.NewPostHandler()
	v1 := r.Group("/article")
	{
		v1.GET("/:limit/:offset", articleHandler.HandleListPost)
		v1.GET("/count", articleHandler.HandleCountStatus)
		v1.POST("", articleHandler.CreateAbleValidator, articleHandler.HandleCreate)
		v1.GET("/detail/:id", articleHandler.HandleReadPost)
		v1.PUT("/:id", articleHandler.UpdateAbleValidator, articleHandler.HandleUpdate)
		v1.DELETE("/:id", articleHandler.HandleDelete)
	}

	r.Run()
}
