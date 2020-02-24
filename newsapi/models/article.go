package models

import (
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

type Article struct {
	Title       string `xml:"title",json:"title"`
	Category    string `xml:"category",json:"category"`
	Link        string `xml:"link",json:"link"`
	Description string `xml:"description",json:"description"`
	PubDate     string `xml:"pubDate",json:"pubDate"`
	Date        time.Time
	UID         string `json:"uid"`
	Source      string `json:"source"`
}

// GetPubDate converts the PubDate from the XML document into Golang-runtime-friendly format.
func (a *Article) GetPubDate() (time.Time, error) {
	t, err := time.Parse(time.RFC1123Z, a.PubDate)
	if err != nil {
		t, err = time.Parse(time.RFC1123, a.PubDate)
		if err != nil {
			return t, err
		}
	}
	//TODO: consider refactoring and adding a wider range of supported formats, just in case
	return t, nil
}

// FetchArticleBody attempts to retrieve the body of the article from its original URL
func (a *Article) FetchArticleBody() ([]byte, error) {
	resp, err := http.Get(a.Link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}

// FillMissingDetails completes Data, Category and UID for the article
func (a *Article) FillMissingDetails() {
	t, err := a.GetPubDate()
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse time provided in article body")
	}
	a.Date = t

	if a.UID == "" {
		a.UID = xid.NewWithTime(a.Date).String()
	}

	if a.Category == "" {
		a.Category = "unspecified"
	}
}
