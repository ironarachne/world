package main

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/ironarachne/world/pkg/cartography"
	"github.com/ironarachne/world/pkg/random"
	"image"
	"image/png"
	"net/http"
	"strconv"
)

func writeImage(w http.ResponseWriter, img *image.Image) error {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		err = fmt.Errorf("failed to encode image: %w", err)
		return err
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		err = fmt.Errorf("failed to write image: %w", err)
		return err
	}

	return nil
}

func getWorldMap(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	im, err := cartography.RenderMap(1024, 1024)
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = writeImage(w, &im)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getWorldMapRandom(w http.ResponseWriter, r *http.Request) {
	im, err := cartography.RenderMap(1024, 1024)
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = writeImage(w, &im)
	if err != nil {
		handleError(w, r, err)
		return
	}
}
