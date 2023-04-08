package service

import (
	"imageOptimisation/entities"
	"imageOptimisation/pkg/repository"

	"github.com/gin-gonic/gin"
)

type ImageOperation interface {
	CreateImage(image entities.Image, filename string) error
	GetImageById(image entities.Image, id int, c *gin.Context) (string, error)
}

type Service struct {
	ImageOperation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ImageOperation: NewImageOperationService(repos.ImageOperation),
	}
}
