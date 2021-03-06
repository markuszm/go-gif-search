package lib

import (
	"errors"
	"github.com/bhaskarsaraogi/giphy"
)

type GiphyClient struct {
	client *giphy.Client
	limit int
}

type Gif struct {
	Id   string
	Name string
	Url  string
}

func NewGiphyClient(apiKey string, limit int) *GiphyClient {
	g := giphy.DefaultClient
	g.APIKey = apiKey
	g.Limit = limit

	return &GiphyClient{client: g, limit: limit}
}

func (g *GiphyClient) SearchGif(keyword string, ranking int) (Gif, error) {
	if ranking >= g.limit {
		return Gif{}, errors.New("ranking is over limit")
	}

	search, err := g.client.Search([]string{keyword})
	if err != nil {
		return Gif{}, err
	}

	if len(search.Data) == 0 {
		return Gif{}, nil
	}

	gif := search.Data[ranking]
	return extractGifFromData(gif), nil
}

func (g *GiphyClient) TranslateGif(phrase string) (Gif, error) {
	translate, err := g.client.Translate([]string{phrase})
	if err != nil {
		return Gif{}, err
	}
	imageData := translate.Data
	return Gif{
		Id:   imageData.ID,
		Name: imageData.Caption,
		Url:  imageData.Images.Original.Mp4,
	}, nil
}

func (g *GiphyClient) TrendingGif(ranking int) (Gif, error) {
	if ranking >= g.limit {
		return Gif{}, errors.New("ranking is over limit")
	}

	trending, err := g.client.Trending()
	if err != nil {
		return Gif{}, err
	}

	gif := trending.Data[ranking]
	return extractGifFromData(gif), nil
}

func (g *GiphyClient) TrendingGifs() ([]Gif, error) {
	trending, err := g.client.Trending()
	if err != nil {
		return []Gif{}, err
	}

	var gifs []Gif
	for _, g := range trending.Data {
		gifs = append(gifs, extractGifFromData(g))
	}

	return gifs, nil
}

func (g *GiphyClient) RandomGif() (Gif, error) {
	random, err := g.client.Random([]string{})
	if err != nil {
		return Gif{}, nil
	}

	imageData := random.Data
	return Gif{
		Id:   imageData.ID,
		Name: imageData.URL,
		Url:  imageData.ImageMp4URL,
	}, nil
}

func extractGifFromData(gifData giphy.Data) Gif {
	img := gifData.Images.FixedHeight
	return Gif{
		Id:   gifData.ID,
		Name: gifData.URL,
		Url:  img.Mp4,
	}
}