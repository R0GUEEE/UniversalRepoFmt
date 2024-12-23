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
	DownloadURL   string               `json:"downloadURL"` // required for sidestore here, for some reason
	Versions      []AltStoreAppVersion `json:"versions"`
	Permissions   UniversalPermissions `json:"appPermissions"` // sidestore has a different permissions format, but it's OK cause it's optional
}

type AltStore struct {
	Name        string        `json:"name"`
	Identifier  string        `json:"identifier"` // required for sidestore
	IconURL     string        `json:"iconURL"`
	Caption     string        `json:"subtitle,omitempty"`
	Description string        `json:"description"`
	Apps        []AltStoreApp `json:"apps"`
}
