package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type ImageService struct {
	UploadDir string
	BaseURL   string
}

func NewImageService(uploadDir, baseURL string) *ImageService {
	return &ImageService{
		UploadDir: uploadDir,
		BaseURL:   baseURL,
	}
}

func (s *ImageService) SaveImage(file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll(s.UploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	ext := filepath.Ext(file.Filename)
	if !isAllowedImageType(ext) {
		return "", fmt.Errorf("unsupported file type: %s", ext)
	}

	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filepath := filepath.Join(s.UploadDir, filename)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return fmt.Sprintf("%s/images/%s", s.BaseURL, filename), nil
}

func isAllowedImageType(ext string) bool {
	ext = strings.ToLower(ext)
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	return allowedTypes[ext]
}