package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"imageOptimisation/entities"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (h *Handler) uploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := fmt.Sprintf("./localStorage/%s", filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	img, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func(img *os.File) {
		err := img.Close()
		if err != nil {

		}
	}(img)
	var image entities.Image
	err = h.services.ImageOperation.CreateImage(image, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

func (h *Handler) getImageById(c *gin.Context) {
	var image entities.Image

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	filename, err := h.services.ImageOperation.GetImageById(image, id, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found"})
		return
	}

	c.File(fmt.Sprintf("./localStorage/%s", filename))
}
