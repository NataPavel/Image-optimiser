package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"imageOptimisation/entities"
)

type ImageOperation interface {
	CreateImage(image entities.Image, filename string) error
	GetImageById(image entities.Image, id int, c *gin.Context) (string, error)
}

type Repository struct {
	ImageOperation
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ImageOperation: NewImageOp(db),
	}
}
