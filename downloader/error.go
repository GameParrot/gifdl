package downloader

import "errors"

var (
	ErrorDownloadLinkNotFound = errors.New("could not find download link in webpage")
)
