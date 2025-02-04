package service

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"mime/multipart"
	"os"
	"strings"

	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ValidateFileIsImage(file *multipart.FileHeader) bool {
	typeFile := strings.ToLower(file.Header.Get("Content-Type"))
	validImageTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}
	return validImageTypes[typeFile]
}

func ConvertToWebP(file *multipart.FileHeader, outputPath string) error {
	// Abrir el archivo
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Decodificar la imagen
	var img image.Image
	typeImage := strings.ToLower(file.Header.Get("Content-Type"))
	fmt.Println("Type image:", typeImage)
	switch typeImage {
	case "image/jpeg":
		img, err = jpeg.Decode(srcFile)
	case "image/png":
		img, err = png.Decode(srcFile)
	default:
		return fmt.Errorf("formato de imagen no soportado")
	}
	if err != nil {
		return err
	}

	fmt.Println("Ruta de salida:", outputPath)

	var buf bytes.Buffer

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 90)
	if err != nil {
		return err
	}

	// Codificar la imagen en formato WebP
	if err := webp.Encode(&buf, img, options); err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, buf.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func ValidateSizeFile(file *multipart.FileHeader, maxSize int64) (bool, domain.InvalidFile) {
	message := fmt.Sprintf("File size is too big: %d, max file exceded (%d)", file.Size, maxSize)
	fmt.Println(message)
	if file.Size > maxSize {
		return false, domain.InvalidFile{
			Name:    file.Filename,
			Message: message,
		}
	}
	return true, domain.InvalidFile{
		Name:    file.Filename,
		Message: "File size is valid",
	}
}

func SerializeFileResponse() {

}

func SaveFile(file multipart.File) {

}
