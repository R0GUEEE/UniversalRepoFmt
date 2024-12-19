package repos

type ScarletApp struct {
	Name          string `json:"name"`
	DeveloperName string `json:"dev"`
	Version       string `json:"version"`
	Icon          string `json:"icon"`
	Down          string `json:"down"`
	Description   string `json:"description"`
	BundleID      string `json:"bundleID"`
}

type ScarletMeta struct {
	RepoName string `json:"repoName"`
	RepoIcon string `json:"repoIcon"`
}

type Scarlet struct {
	Meta ScarletMeta  `json:"META"`
	Apps []ScarletApp `json:"Tweaked"`
}
