/*
Package save implements a system for saving files to various targets
*/
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

// ToDOSpaces saves a file to Digital Ocean spaces with the given MIME type
func ToDOSpaces(filePath string, reader io.Reader, contentType string) (string, error) {
	accessKey := os.Getenv("SPACES_KEY")
	secretKey := os.Getenv("SPACES_SECRET")
	bucketName := os.Getenv("SPACES_BUCKET_NAME")
	domainName := os.Getenv("STATIC_DOMAIN_NAME")
	endpoint := "nyc3.digitaloceanspaces.com"

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		err = fmt.Errorf("failed to connect to Digital Ocean Spaces: %w", err)
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
		err = fmt.Errorf("failed to save to Digital Ocean Spaces: %w", err)
		return "", err
	}

	url := "https://" + domainName + "/" + filePath

	return url, nil
}

// PNG saves a PNG image to permanent storage. It relies on an environment variable WORLDAPI_SAVE_TARGET to determine behavior.
// The default behavior is to save to the local filesystem. Set WORLDAPI_SAVE_DIRECTORY to the directory to use as the root directory
// for saving files. Set WORLDAPI_WEB_DOMAIN to the domain name (without protocol) the API serves on.
func PNG(directory string, fileName string, img image.Image) (string, error) {
	var result string
	var err error

	saveTarget := os.Getenv("WORLDAPI_SAVE_TARGET")
	filePath := directory + "/" + fileName

	if saveTarget == "DO" {
		result, err = PNGToDO(filePath, img)
		if err != nil {
			err = fmt.Errorf("failed to save PNG: %w", err)
			return "", err
		}
		return result, nil
	}

	saveDirectory := os.Getenv("WORLDAPI_SAVE_DIRECTORY")
	savePath := saveDirectory + "/images/" + fileName
	fg, err := os.Create(savePath)
	defer fg.Close()
	if err != nil {
		err = fmt.Errorf("failed to save PNG: %w", err)
		return "", err
	}
	err = png.Encode(fg, img)
	if err != nil {
		err = fmt.Errorf("failed to save PNG: %w", err)
		return "", err
	}

	webDomain := os.Getenv("WORLDAPI_WEB_DOMAIN")
	url := "https://" + webDomain + "/" + filePath

	return url, nil
}

// PNGToDO saves a PNG image to the given space
func PNGToDO(filePath string, img image.Image) (string, error) {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		err := png.Encode(writer, img)
		if err != nil {
			panic(err)
		}
	}()

	url, err := ToDOSpaces(filePath, reader, "image/png")
	if err != nil {
		err = fmt.Errorf("Failed to save PNG to Digital Ocean Spaces: %w", err)
		return "", err
	}

	return url, nil
}

// SVGToDO saves an SVG image to the given space
func SVGToDO(filePath string, contents string) (string, error) {
	reader := strings.NewReader(contents)

	url, err := ToDOSpaces(filePath, reader, "image/svg+xml")
	if err != nil {
		err = fmt.Errorf("Failed to save SVG to Digital Ocean Spaces: %w", err)
		return "", err
	}

	return url, nil
}
