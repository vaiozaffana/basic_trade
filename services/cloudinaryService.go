package services

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func InitCloudinary() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}
	return cld, nil
}

func UploadToCloudinary(filepath string, fileName string) (string, error) {
	cld, err := InitCloudinary()
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		PublicID:       fileName,
		Folder:         "product_images",
		AllowedFormats: []string{"jpg", "jpeg", "png", "svg"},
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), filepath, uploadParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to Cloudinary: %v", err)
	}

	return uploadResult.SecureURL, nil
}
