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

func ship(filePath string, reader io.Reader, contentType string) string {
	accessKey := os.Getenv("SPACES_KEY")
	secretKey := os.Getenv("SPACES_SECRET")
	bucketName := os.Getenv("SPACES_BUCKET_NAME")
	endpoint := "nyc3.digitaloceanspaces.com"
	useSSL := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		panic(err) // TODO: Adopt Go 1.12 fmt.Errorf pattern
	}

	userMetadata := make(map[string]string)
	userMetadata["x-amz-acl"] = "public-read"

	opts := minio.PutObjectOptions{
		ContentType: contentType,
		UserMetadata: userMetadata,
	}
	n, err := client.PutObject(bucketName, filePath, reader, -1, opts)
	if err != nil {
		panic(err) // TODO: Adopt Go 1.12 fmt.Errorf pattern
	}
	fmt.Println(n)

	url := "https://" + bucketName + "." + endpoint + "/" + filePath

	return url
}

// PNG saves a PNG image to the given space
func PNG(filePath string, img image.Image) string {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		err := png.Encode(writer, img)
		if err != nil {
			panic(err) // TODO: Adopt Go 1.12 fmt.Errorf pattern
		}
	}()

	url := ship(filePath, reader, "image/png")

	return url
}

// SVG saves an SVG image to the given space
func SVG(filePath string, contents string) string {
	reader := strings.NewReader(contents)

	url := ship(filePath, reader, "image/svg+xml")

	return url
}
