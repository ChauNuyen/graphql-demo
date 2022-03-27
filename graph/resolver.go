package graph

import "github.com/chaunuyen/grapql-demo/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	videos []*model.Video
}
