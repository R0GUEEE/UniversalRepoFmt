package repos

type AltStoreAppVersion struct {
	Version     string `json:"version"`
	Date        string `json:"date"`
	Size        int64  `json:"size"`
	DownloadURL string `json:"downloadURL"`
	Description string `json:"localizedDescription"`
}

type AltStoreApp struct {
	Name          string               `json:"name"`
	DeveloperName string               `json:"developerName"`
	BundleID      string               `json:"bundleIdentifier"`
	Caption       string               `json:"subtitle,omitempty"`
	Description   string               `json:"localizedDescription"`
	IconURL       string               `json:"iconURL"`
	Versions      []AltStoreAppVersion `json:"versions"`
}

type AltStore struct {
	Name        string        `json:"name"`
	Caption     string        `json:"subtitle,omitempty"`
	Description string        `json:"description"`
	Apps        []AltStoreApp `json:"apps"`
}
