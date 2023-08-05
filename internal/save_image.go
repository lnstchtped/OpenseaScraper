package internal

import (
	"fmt"
	"os"
)

func SaveImage(imageBytes []byte, collection, filename string) error {
	outDir := fmt.Sprintf("scraped/%s", collection)
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s", outDir, filename), imageBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
