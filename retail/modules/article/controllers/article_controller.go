package controllers

import (
	"net/http"
	"strconv"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	// Get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// Find the article from the database
	article, err := controller.articleService.Find(id)
}
