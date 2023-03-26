package service

import (
	"github.com/gin-gonic/gin"
	"imageOptimisation/entities"
	"imageOptimisation/pkg/repository"
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
