package downloader

import (
	"net/url"
	"strings"
)

type GiphyGifDownloader struct {
}

func (g *GiphyGifDownloader) DownloadURL(url string) (string, string, error) {
	downloadUrl, title, err := metaDownload(url, "property", "og:image")
	if err != nil {
		return "", "", err
	}
	if strings.HasSuffix(downloadUrl, ".webp") {
		downloadUrl = strings.TrimSuffix(downloadUrl, "webp") + "gif"
	}
	splitTitle := strings.Split(title, " GIF by ")
	if len(splitTitle) > 1 {
		title = splitTitle[0]
	}
	return downloadUrl, title, nil
}

func (g *GiphyGifDownloader) MatchesURL(rawUrl string) bool {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return false
	}
	host := strings.ToLower(url.Host)
	return host == "giphy.com" || host == "www.giphy.com"
}
