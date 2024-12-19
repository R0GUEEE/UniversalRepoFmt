package repos

type ESignApp struct {
	Name        string `json:"name"`
	BundleID    string `json:"bundleIdentifier"`
	Description string `json:"localizedDescription"`
	DownloadURL string `json:"downloadURL"`
	IconURL     string `json:"iconURL"`
	Version     string `json:"version"`
	Date        string `json:"versionDate"`
	Size        int64  `json:"size"`
}

type ESign struct {
	Name       string     `json:"name"`
	Identifier string     `json:"identifier,omitempty"`
	Apps       []ESignApp `json:"apps"`
}
