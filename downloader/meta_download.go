package downloader

import (
	"fmt"
	"net/http"

	"github.com/gameparrot/gifdl/utils"
	"golang.org/x/net/html"
)

func metaDownload(url string, key string, val string) (string, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:133.0) Gecko/20100101 Firefox/133.0")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("download webpage: %w", err)
	}
	htmlBodyResp := res.Body
	doc, err := html.Parse(htmlBodyResp)
	htmlBodyResp.Close()
	if err != nil {
		return "", "", fmt.Errorf("parse html: %w", err)
	}
	downloadUrl := utils.FindMeta(doc, val, key)
	if downloadUrl == "" {
		return "", "", ErrorDownloadLinkNotFound
	}
	return downloadUrl, utils.GetTitle(doc), nil
}
