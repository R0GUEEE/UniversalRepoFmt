package UniversalRepoFmt

import "github.com/asdfzxcvbn/UniversalRepoFmt/repos"

func ConvertToESign(uni *repos.Universal) *repos.ESign {
	r := repos.ESign{
		Name:       uni.Name,
		Identifier: uni.Identifier,
	}

	for _, app := range uni.Apps {
		r.Apps = append(r.Apps, repos.ESignApp{
			Name:        app.Name,
			BundleID:    app.BundleID,
			Description: app.Description,
			DownloadURL: app.DownloadURL,
			IconURL:     app.IconURL,
			Version:     app.Version,
			Date:        app.Date,
			Size:        app.Size,
		})
	}

	return &r
}

func ConvertToGBox(uni *repos.Universal) *repos.GBox {
	r := repos.GBox{
		Version:     "1.0",
		Name:        uni.Name,
		IconURL:     uni.IconURL,
		Description: uni.Description,
	}

	for _, app := range uni.Apps {
		r.Apps = append(r.Apps, repos.GBoxApp{
			Type:        "SELF_SIGN", // wtf?
			Date:        app.Date,
			Name:        app.Name,
			Version:     app.Version,
			Description: app.Description,
			IconURL:     app.IconURL,
			DownloadURL: app.DownloadURL,
		})
	}

	return &r
}

func ConvertToFeather(uni *repos.Universal) *repos.Feather {
	r := repos.Feather{
		Name:       uni.Name,
		Identifier: uni.Identifier,
	}

	for _, app := range uni.Apps {
		r.Apps = append(r.Apps, repos.FeatherApp{
			Name:        app.Name,
			BundleID:    app.BundleID,
			Caption:     app.Caption,
			Description: app.Description,
			DownloadURL: app.DownloadURL,
			IconURL:     app.IconURL,
			Version:     app.Version,
			Size:        app.Size,
		})
	}

	return &r
}

func ConvertToAltStore(uni *repos.Universal) *repos.AltStore {
	r := repos.AltStore{
		Name:        uni.Name,
		Identifier:  uni.Identifier,
		IconURL:     uni.IconURL,
		Caption:     uni.Caption,
		Description: uni.Description,
	}

	// step 1: find indexes of all apps with same bundle id
	bundleIdIdxs := make(map[string][]int, len(uni.Apps))
	for idx, app := range uni.Apps {
		_, ok := bundleIdIdxs[app.BundleID]
		if !ok {
			bundleIdIdxs[app.BundleID] = []int{idx}
		} else {
			bundleIdIdxs[app.BundleID] = append(bundleIdIdxs[app.BundleID], idx)
		}
	}

	for bundle, idxs := range bundleIdIdxs {
		orig := uni.Apps[idxs[0]]
		app := repos.AltStoreApp{
			Name:          orig.Name,
			DeveloperName: orig.DeveloperName,
			BundleID:      bundle,
			IconURL:       orig.IconURL,
		}

		if permissions, ok := uni.Permissions[bundle]; ok {
			app.Permissions = permissions
		}

		for _, idx := range idxs {
			origIdx := uni.Apps[idx]
			app.Versions = append(app.Versions, repos.AltStoreAppVersion{
				Version:     origIdx.Version,
				Date:        origIdx.Date,
				Description: origIdx.Description,
				DownloadURL: origIdx.DownloadURL,
				Size:        origIdx.Size,
			})
		}

		r.Apps = append(r.Apps, app)
	}

	return &r
}

func ConvertToScarlet(uni *repos.Universal) *repos.Scarlet {
	r := repos.Scarlet{
		Meta: repos.ScarletMeta{
			RepoName: uni.Name,
			RepoIcon: uni.IconURL,
		},
	}

	for _, app := range uni.Apps {
		r.Apps = append(r.Apps, repos.ScarletApp{
			Name:          app.Name,
			DeveloperName: app.DeveloperName,
			Version:       app.Version,
			Icon:          app.IconURL,
			Down:          app.DownloadURL,
			Description:   app.Description,
			BundleID:      app.BundleID,
		})
	}

	return &r
}

// a convenience function that returns all formatted repos.
func ConvertToAll(uni *repos.Universal) *repos.All {
	return &repos.All{
		ESign:    ConvertToESign(uni),
		GBox:     ConvertToGBox(uni),
		Feather:  ConvertToFeather(uni),
		AltStore: ConvertToAltStore(uni),
		Scarlet:  ConvertToScarlet(uni),
	}
}
