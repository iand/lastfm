/*
  PUBLIC DOMAIN STATEMENT
  To the extent possible under law, Ian Davis has waived all copyright
  and related or neighboring rights to this Source Code file.
  This work is published from the United Kingdom.
*/

// Client for the Last.fm API
package lastfm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey string
}

type Artist struct {
	Name       string           `json:"name"`
	MBID       string           `json:"mbid"`
	URL        string           `json:"url"`
	Image      []Image          `json:"image"`
	Stats      ArtistStats      `json:"stats"`
	Streamable string           `json:"streamable"`
	OnTour     string           `json:"ontour"`
	Similar    SimilarContainer `json:"similar"`
	Tags       TagContainer     `json:"tags"`
	Bio        ArtistBio        `json:"bio"`
}

type Image struct {
	URL  string `json:"#text"`
	Size string `json:"size"`
}

type ArtistStats struct {
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
}
type SimilarContainer struct {
	Artist []Artist `'json:"artist"`
}

type TagContainer struct {
	Tag []Tag `'json:"tag"`
}

type Tag struct {
	Name string `'json:"name"`
	URL  string `'json:"url"`
}

type WikiInfo struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type ArtistBio struct {
	Links      LinkContainer `'json:"links"`
	YearFormed string        `json:"yearformed"`
	WikiInfo
}

type LinkContainer struct {
	Link Link `'json:"link"`
}

type Link struct {
	Text string `json:"#text"`
	Rel  string `json:"rel"`
	URL  string `json:"href"`
}

type Track struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	MBID       string       `json:"mbid"`
	URL        string       `json:"url"`
	Duration   string       `json:"duration"`
	Listeners  string       `json:"listeners"`
	Playcount  string       `json:"playcount"`
	Streamable Streamable   `json:"streamable"`
	Artist     Artist       `json:"artist"`
	Album      Album        `json:"album"`
	TopTags    TagContainer `json:"toptags"`
	Wiki       WikiInfo     `json:"wiki"`
}

type Streamable struct {
	Text      string `json:"#text"`
	Fulltrack string `json:"fulltrack"`
}

type Album struct {
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	MBID   string  `json:"mbid"`
	URL    string  `json:"url"`
	Image  []Image `json:"image"`
}

func New(apikey string) *Client {
	return &Client{APIKey: apikey}
}

func (client *Client) ArtistInfoByName(artist string, lang string, username string) (*Artist, error) {
	var data struct {
		Artist Artist `json:"artist"`
	}

	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.getInfo&format=json&artist=%s&api_key=%s", url.QueryEscape(artist), url.QueryEscape(client.APIKey))
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("ArtistInfoByName failed with http error: %s", err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("ArtistInfoByName failed to parse JSON response: %s", err.Error())
	}

	if err != nil {
		return nil, err
	}
	return &data.Artist, nil
}

func (client *Client) TrackInfoByName(track string, artist string, username string) (*Track, error) {
	var data struct {
		Track Track `json:"track"`
	}

	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=track.getInfo&format=json&track=%s&artist=%s&api_key=%s", url.QueryEscape(track), url.QueryEscape(artist), url.QueryEscape(client.APIKey))
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("TrackInfoByName failed with http error: %s", err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("TrackInfoByName failed to parse JSON response: %s", err.Error())
	}

	if err != nil {
		return nil, err
	}
	return &data.Track, nil
}
