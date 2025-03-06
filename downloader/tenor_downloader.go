package downloader

import (
	"net/url"
	"strings"
)

type TenorGifDownloader struct {
}

func (t *TenorGifDownloader) DownloadURL(url string) (string, string, error) {
	downloadUrl, title, err := metaDownload(url, "property", "og:image")
	if err != nil {
		return "", "", err
	}
	if strings.HasSuffix(downloadUrl, ".webp") {
		downloadUrl = strings.TrimSuffix(downloadUrl, "webp") + "gif"
	}
	splitTitle := strings.Split(title, " - ")
	if len(splitTitle) > 1 {
		title = splitTitle[1]
	}
	return downloadUrl, title, nil
}

func (t *TenorGifDownloader) MatchesURL(rawUrl string) bool {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return false
	}
	host := strings.ToLower(url.Host)
	return host == "tenor.com" || host == "www.tenor.com"
}
