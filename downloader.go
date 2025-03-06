package gifdl

type Downloader interface {
	// DownloadURL gets the download url and title of a gif
	DownloadURL(url string) (downloadUrl string, title string, err error)
	// MatchesURL checks if a url is handled by this downloader
	MatchesURL(url string) bool
}
