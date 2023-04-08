package repository

import (
	"database/sql"
	"imageOptimisation/entities"

	"github.com/gin-gonic/gin"
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
