# UniversalRepoFmt
a universal IPA repo format that automatically converts into nearly every major format! it converts into formats for the following apps:

- AltStore (incl. SideStore, altsource-viewer, and derivatives)
- ESign (can usually be used in most places)
- Feather (incl. TrollApps)
- GBox
- Scarlet

[![Go Reference](https://pkg.go.dev/badge/github.com/asdfzxcvbn/UniversalRepoFmt.svg)](https://pkg.go.dev/github.com/asdfzxcvbn/UniversalRepoFmt)

## format
most (if not all) keys are technically optional. here's an example repo with every possible key:

```json
{
  "name": "hello!",
  "identifier": "fyi.zxcvbn.repo.test",
  "iconURL": "https://example.com/path/to/some/icon.png",
  "caption": "a *short* description of your repo!",
  "description": "this can be a longer description and include stuff like your links or something, idk",
  "apps": [
    {
      "name": "app name",
      "developerName": "zx",
      "bundleID": "fyi.zxcvbn.app",
      "caption": "just like the repo's caption",
      "description": "just like the repo's description",
      "downloadURL": "https://example.com/path/to/some/app.ipa",
      "iconURL": "https://example.com/hopefully/path/to/itunes/icon.png",
      "version": "1.2",
      "date": "2024-12-18",
      "size": 1073741824
    }
  ],
  "permissions": {
    "fyi.zxcvbn.app": {
      "entitlements": [
        "aps-environment",
        "com.apple.developer.associated-domains",
        "keychain-access-groups",
        "com.apple.security.application-groups"
      ],
      "privacy": {
        "NSBluetoothAlwaysUsageDescription": "example",
        "NSFaceIDUsageDescription": "for face id"
      }
    }
  }
}
```

you can see the output for this example in the `examples/` directory.

## useful key explanations
### root
`identifier` is usually a reverse id. this can really be anything, but i like the reverse domains.

`caption` translates to the `subtitle` key in the AltStore and Feather outputs. it's the tiny text shown under the repo name.

`description` is the big text describing your repo. this is only shown in some apps so feel free to exclude it.

`permissions` translates to the `appPermissions` key in the AltStore format. this is required according to my memory and [AltStore docs](https://faq.altstore.io/developers/make-a-source#apppermissions-app-permissions-object). if you're going to upload multiple versions of an app with the same bundle id, chances are their permissions are the same too, which is why `permissions` is a map of bundle ids to permissions.

### app
`developerName` is only shown in some apps. exclude it if you want.

`bundleID` is the bundle identifier of the ipa. you want this to be accurate, it'll be used to make the AltStore-formatted repo accurate.

`caption` and `description` basically have the same purpose as their root counterparts.

`date` is the date which this ipa was uploaded. this key isnt required, but can be helpful, especially in some apps like gbox. the most common date format is `2024-12-18`, however this doesn't work in gbox. gbox uses the format `2024-12-18T00:00:00+0800`. you can pick either.

## usage
### as a cli tool
you can either download the binary from the latest release or build this yourself with the latest version of [Go](https://go.dev):

```bash
$ cd cmd/urf
$ go build -ldflags="-w -s"
```

```bash
$ urf -h
Usage of urf:
  -input string
    	the universal-format json file to read from (default "universal.json")
  -output string
    	the directory to write the output repos to (default "formatted")
```

### as a library
see the go reference at the top of the readme. just import `repos`, construct a `Universal` struct, then pass it to any of the ConvertTo functions. by the way, this tool is under the Unlicense, feel free to use it anywhere! credit isn't necessary per the license, but i would appreciate it !
