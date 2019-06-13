package controllers

import (
	"github.com/OahcUil94/go-training/iris-training/overview/repositories"
	"github.com/OahcUil94/go-training/iris-training/overview/services"
	"github.com/kataras/iris/mvc"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)

	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}