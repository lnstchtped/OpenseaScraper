package pkg

import (
	"github.com/lnstchtped/OpenseaScraper/internal"
	"sync"
)

func RunDownloader(slug string, proxy string) error {
	s, err := NewSession(proxy)
	if err != nil {
		return err
	}

	c, scraped, err := s.ScrapeImages(slug, "")
	internal.DownloadImages(slug, scraped)
	for c != "" && err == nil {
		c, scraped, err = s.ScrapeImages(slug, c)
		if len(scraped) == 0 {
			break
		}
		var wg sync.WaitGroup
		wg.Add(len(scraped))
		for _, i := range scraped {
			go func(s string) {
				defer wg.Done()
				internal.DownloadImages(slug, []string{s})
			}(i)
		}
		wg.Wait()
	}

	return err
}
