package repos

type GBoxApp struct {
	Type        string `json:"appType"`
	Date        string `json:"versionDate"` // changing from "appUpdateTime" so it shows updated time correctly with yyyy-mm-dd date format (common)
	Name        string `json:"appName"`
	Version     string `json:"appVersion"`
	Description string `json:"appDescription"`
	IconURL     string `json:"appImage"`
	DownloadURL string `json:"appPackage"`
}

type GBox struct {
	Version     string    `json:"version"`
	Name        string    `json:"sourceName"`
	IconURL     string    `json:"sourceImage"`
	Description string    `json:"sourceDescription,omitempty"`
	Apps        []GBoxApp `json:"appRepositories"`
}
