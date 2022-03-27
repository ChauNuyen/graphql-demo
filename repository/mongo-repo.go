package repository

import (
	"github.com/chaunuyen/grapql-demo/graph/model"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []model.Video
}

type database struct {
	client *mongo.Client
}

func New() VideoRepository {
	return &database{
		client: nil,
	}
}
