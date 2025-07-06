package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title          string    `xml:"title"`
		Description    string    `xml:"description"`
		Link           string    `xml:"link"`
		Category       string    `xml:"category"`
		Copyright      string    `xml:"copyright"`
		Docs           string    `xml:"docs"`
		Language       string    `xml:"language"`
		LastBuildDate  string    `xml:"lastBuildDate"`
		ManagingEditor string    `xml:"managingEditor"`
		PubDate        string    `xml:"pubDate"`
		WebMaster      string    `xml:"webMaster"`
		Generator      string    `xml:"generator"`
		Item           []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	Enclosure   string `xml:"enclosure"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Source      string `xml:"source"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssFeed, nil
}

// func demoRSS(url string) (string, error) {
// 	httpClient := http.Client{
// 		Timeout: 10 * time.Second,
// 	}

// 	resp, err := httpClient.Get(url)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(data), nil

// }
