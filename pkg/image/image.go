package image

import (
	"context"
	"log"
)

// Image ...
type Image interface {
	Process(ctx context.Context)
}

type image struct {
}

// Process ...
func (i *image) Process(ctx context.Context) {
	log.Println("image processing")
}

// NewImage ...
func NewImage() Image {
	return &image{}
}
