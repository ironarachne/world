package save

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/minio/minio-go"
)

func ship(filePath string, reader io.Reader, contentType string) (string, error) {
	accessKey := os.Getenv("SPACES_KEY")
	secretKey := os.Getenv("SPACES_SECRET")
	bucketName := os.Getenv("SPACES_BUCKET_NAME")
	domainName := os.Getenv("STATIC_DOMAIN_NAME")
	endpoint := "nyc3.digitaloceanspaces.com"

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		err = fmt.Errorf("Failed to save to Digital Ocean Spaces: %w", err)
		return "", err
	}

	userMetadata := make(map[string]string)
	userMetadata["x-amz-acl"] = "public-read"

	opts := minio.PutObjectOptions{
		ContentType:  contentType,
		UserMetadata: userMetadata,
	}
	_, err = client.PutObject(bucketName, filePath, reader, -1, opts)
	if err != nil {
		err = fmt.Errorf("Failed to save to Digital Ocean Spaces: %w", err)
		return "", err
	}

	url := "https://" + domainName + "/" + filePath

	return url, nil
}

// PNG saves a PNG image to the given space
func PNG(filePath string, img image.Image) (string, error) {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		err := png.Encode(writer, img)
		if err != nil {
			panic(err)
		}
	}()

	url, err := ship(filePath, reader, "image/png")
	if err != nil {
		err = fmt.Errorf("Failed to save PNG to Digital Ocean Spaces: %w", err)
		return "", err
	}

	return url, nil
}

// SVG saves an SVG image to the given space
func SVG(filePath string, contents string) (string, error) {
	reader := strings.NewReader(contents)

	url, err := ship(filePath, reader, "image/svg+xml")
	if err != nil {
		err = fmt.Errorf("Failed to save SVG to Digital Ocean Spaces: %w", err)
		return "", err
	}

	return url, nil
}
