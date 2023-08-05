package internal

import "log"

func DownloadImages(col string, urls []string) {
	for _, url := range urls {
		img, fil, err := DownloadImage(url)
		if err != nil {
			log.Printf("failed to download image: %v", err)
			continue
		}
		if img == nil {
			continue
		}

		err = SaveImage(img, col, fil)
		if err != nil {
			log.Printf("failed to save image: %v", err)
			continue
		}

		log.Printf("saved image: %s", fil)
	}
}
