package service

import (
	"github.com/gin-gonic/gin"
	"imageOptimisation/entities"
	"imageOptimisation/pkg/repository"
)

type ImageOperationService struct {
	repos repository.ImageOperation
}

func NewImageOperationService(repos repository.ImageOperation) *ImageOperationService {
	return &ImageOperationService{repos: repos}
}

func (s *ImageOperationService) CreateImage(image entities.Image, filename string) error {
	return s.repos.CreateImage(image, filename)
}

func (s *ImageOperationService) GetImageById(image entities.Image, id int, c *gin.Context) (string, error) {
	return s.repos.GetImageById(image, id, c)
}
