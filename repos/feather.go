package repos

type FeatherApp struct {
	Name        string `json:"name"`
	BundleID    string `json:"bundleIdentifier"`
	Caption     string `json:"subtitle,omitempty"`
	Description string `json:"localizedDescription"`
	DownloadURL string `json:"downloadURL"`
	IconURL     string `json:"iconURL"`
	Version     string `json:"version"`
	Size        int64  `json:"size"`
}

type Feather struct {
	Name       string       `json:"name"`
	Identifier string       `json:"identifier,omitempty"`
	Apps       []FeatherApp `json:"apps"`
}
