package gifdl

import (
	"errors"
	"fmt"
	"image/gif"
	"net/http"

	"github.com/gameparrot/gifdl/downloader"
)

var ErrUnsupportedURL = errors.New("no downloader for url")

var downloaders = map[string]Downloader{}

// RegisterDownloader registers a GIF downloader for a URL pattern
func RegisterDownloader(name string, downloader Downloader) {
	downloaders[name] = downloader
}

func init() {
	RegisterDownloader("tenor", &downloader.TenorGifDownloader{})
	RegisterDownloader("giphy", &downloader.GiphyGifDownloader{})
}

// DownloadGIF fetches the GIF and decodes it to a *gif.GIF
func DownloadGIF(url string) (*gif.GIF, string, error) {
	downloadUrl, title, err := GetGIFDownloadUrl(url)
	if err != nil {
		return nil, "", err
	}

	resp, err := http.Get(downloadUrl)
	if err != nil {
		return nil, "", fmt.Errorf("download gif: %w", err)
	}

	defer resp.Body.Close()
	gif, err := gif.DecodeAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("read gif: %w", err)
	}

	return gif, title, nil
}

// GetGIFDownloadUrl returns the download URL and title of a GIF
func GetGIFDownloadUrl(url string) (string, string, error) {
	for _, dl := range downloaders {
		if dl.MatchesURL(url) {
			return dl.DownloadURL(url)
		}
	}
	return "", "", ErrUnsupportedURL
}
