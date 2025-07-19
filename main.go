package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/cors"
)

type WebsiteDetails struct {
	Title           string   `json:"title"`
	MetaDescription string   `json:"meta_description"`
	Headings        []string `json:"headings"`
	Paragraphs      []string `json:"paragraphs"`
	Links           []string `json:"links"`
}

func fetchWebsiteDetails(url string) (*WebsiteDetails, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the website: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the HTML: %v", err)
	}

	var metaDescription string
	doc.Find("meta[name='description']").Each(func(i int, s *goquery.Selection) {
		metaDescription, _ = s.Attr("content")
	})

	var headings []string
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		headings = append(headings, s.Text())
	})

	var paragraphs []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			links = append(links, link)
		}
	})

	details := &WebsiteDetails{
		Title:           doc.Find("title").Text(),
		MetaDescription: metaDescription,
		Headings:        headings,
		Paragraphs:      paragraphs,
		Links:           links,
	}

	return details, nil
}

func handleAiravataExtension(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.URL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	details, err := fetchWebsiteDetails(requestData.URL)
	if err != nil {
		http.Error(w, "Failed to fetch website details: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(details)
}

func main() {
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(http.HandlerFunc(handleAiravataExtension))

	http.Handle("/fetch", handler)
	log.Println("Airavata Scrapper is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
