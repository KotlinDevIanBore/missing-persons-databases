package service

import (
    "context"
    "fmt"
    "mime/multipart"
    "path/filepath"
    "strings"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/google/uuid"
)

type ImageService struct {
    S3Client   *s3.Client
    BucketName string
    BaseURL    string
}

func NewImageService(bucketName, baseURL string) (*ImageService, error) {
    if bucketName == "" {
        return nil, fmt.Errorf("bucket name is required")
    }

    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        return nil, fmt.Errorf("unable to load SDK config: %w", err)
    }

    client := s3.NewFromConfig(cfg)
    
    return &ImageService{
        S3Client:   client,
        BucketName: bucketName,
        BaseURL:    baseURL,
    }, nil
}

func (s *ImageService) SaveImage(file *multipart.FileHeader) (string, error) {
    if file == nil {
        return "", fmt.Errorf("no file provided")
    }

    ext := filepath.Ext(file.Filename)
    if !isAllowedImageType(ext) {
        return "", fmt.Errorf("unsupported file type: %s", ext)
    }

    filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
    
    src, err := file.Open()
    if err != nil {
        return "", fmt.Errorf("failed to open uploaded file: %w", err)
    }
    defer src.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    _, err = s.S3Client.PutObject(ctx, &s3.PutObjectInput{
        Bucket: aws.String(s.BucketName),
        Key:    aws.String(filename),
        Body:   src,
        ContentType: aws.String(getContentType(ext)),
        Metadata: map[string]string{
            "OriginalFilename": file.Filename,
            "UploadTimestamp": time.Now().UTC().Format(time.RFC3339),
        },
    })
    if err != nil {
        return "", fmt.Errorf("failed to upload to S3: %w", err)
    }

    return s.GetImageURL(filename), nil
}

func (s *ImageService) GetImageURL(filename string) string {
    return fmt.Sprintf("%s/%s", s.BaseURL, filename)
}

func (s *ImageService) ValidateImageURL(url string) bool {
    return strings.HasPrefix(url, s.BaseURL)
}

func (s *ImageService) DeleteImage(filename string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := s.S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
        Bucket: aws.String(s.BucketName),
        Key:    aws.String(filename),
    })
    return err
}

func getContentType(ext string) string {
    switch strings.ToLower(ext) {
    case ".jpg", ".jpeg":
        return "image/jpeg"
    case ".png":
        return "image/png"
    case ".gif":
        return "image/gif"
    default:
        return "application/octet-stream"
    }
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