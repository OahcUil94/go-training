package repositories

import "github.com/OahcUil94/go-training/iris-training/overview/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name: "imooc video"}
	return movie.Name
}