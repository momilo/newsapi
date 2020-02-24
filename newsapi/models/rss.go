package models

import (
	"bytes"
	"encoding/xml"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type rssFeed struct {
	Version     string    `xml:"version,attr"`
	Title       string    `xml:"channel>title"`
	Link        string    `xml:"channel>link"`
	Description string    `xml:"channel>description"`
	Articles    []Article `xml:"channel>item"`
}

// FetchArticles fetches articles details from an RSS feed url and returns a slice of these articles
func FetchArticles(rssUrl string) []Article {
	//TODO: check for url validity
	response, err := http.Get(rssUrl)
	if err != nil {
		log.Error().Err(err).Str("url", rssUrl).Msg("Error in retrieving rss feed")
		return nil
	}
	defer response.Body.Close()

	xmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Str("url", rssUrl).Msg("Error in reading xml file")
		return nil
	}
	rss := new(rssFeed)
	decoded := xml.NewDecoder(bytes.NewBuffer(xmlData))
	if err = decoded.Decode(rss); err != nil {
		log.Error().Err(err).Str("url", rssUrl).Msg("Error in decoding rss data")
		return nil
	}

	return rss.Articles
}
