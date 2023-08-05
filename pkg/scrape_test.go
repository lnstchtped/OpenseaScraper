package pkg

import "testing"

func TestScrapeImages(t *testing.T) {
	s, err := NewSession("")
	if err != nil {
		t.Fatal(err)
	}

	c, scraped, err := s.ScrapeImages("milady", "")
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range scraped {
		t.Log(s)
	}
	for c != "" {
		c, scraped, err = s.ScrapeImages("milady", c)
		if err != nil {
			t.Fatal(err)
		}
		for _, s := range scraped {
			t.Log(s)
		}
	}

	t.Log("Done")

}
