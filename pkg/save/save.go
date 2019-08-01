package save

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go"
)

func ship(filePath string, reader io.Reader, contentType string) {
	accessKey := os.Getenv("SPACES_KEY")
	secretKey := os.Getenv("SPACES_SECRET")
	endpoint := "nyc3.digitaloceanspaces.com"
	useSSL := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		log.Fatal(err)
	}

	n, err := client.PutObject("ia-assets", filePath, reader, -1, minio.PutObjectOptions{ContentType: contentType})

	fmt.Println("Uploaded image "+filePath+" with bytes ", n)
}

// PNG saves a PNG image to the given space
func PNG(filePath string, img image.Image) {
	reader, writer := io.Pipe()
	png.Encode(writer, img)
	ship(filePath, reader, "image/png")
}

// SVG saves an SVG image to the given space
func SVG(filePath string, contents string) {
	reader := strings.NewReader(contents)
	ship(filePath, reader, "image/svg+xml")
}
