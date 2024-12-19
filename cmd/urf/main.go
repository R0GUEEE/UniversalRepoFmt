package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/asdfzxcvbn/UniversalRepoFmt"
	"github.com/asdfzxcvbn/UniversalRepoFmt/repos"
)

const FilePerm = 0755 // a+rwx,g-w,o-w

var (
	inputFile string
	outputDir string
)

func init() {
	flag.StringVar(&inputFile, "input", "universal.json", "the universal-format json file to read from")
	flag.StringVar(&outputDir, "output", "formatted", "the directory to write the output repos to")
	flag.Parse()
}

// kinda polluted namespace ig, but it's OK!
func createFile(uni *repos.Universal, file string, format FormatType) {
	var formatted any

	switch format {
	case ESign:
		file = strings.Replace(file, ".json", ".esign.json", 1)
		formatted = UniversalRepoFmt.ConvertToESign(uni)
	case GBox:
		file = strings.Replace(file, ".json", ".gbox.json", 1)
		formatted = UniversalRepoFmt.ConvertToGBox(uni)
	case Feather:
		file = strings.Replace(file, ".json", ".feather.json", 1)
		formatted = UniversalRepoFmt.ConvertToFeather(uni)
	case AltStore:
		file = strings.Replace(file, ".json", ".altstore.json", 1)
		formatted = UniversalRepoFmt.ConvertToAltStore(uni)
	case Scarlet:
		file = strings.Replace(file, ".json", ".scarlet.json", 1)
		formatted = UniversalRepoFmt.ConvertToScarlet(uni)
	}

	file = filepath.Join(outputDir, file)

	final, _ := json.MarshalIndent(formatted, "", "  ")
	if err := os.WriteFile(file, final, FilePerm); err != nil {
		log.Fatalf("couldn't write %s: %v\n", file, err)
	}

	log.Println("created " + file)
}

func main() {
	contents, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var uni *repos.Universal
	if err = json.Unmarshal(contents, &uni); err != nil {
		log.Fatal(err)
	}

	if err = os.MkdirAll(outputDir, FilePerm); err != nil {
		log.Fatalf("couldn't make output dir: %v\n", err)
	}

	basename := filepath.Base(inputFile)
	createFile(uni, basename, ESign)
	createFile(uni, basename, GBox)
	createFile(uni, basename, Feather)
	createFile(uni, basename, AltStore)
	createFile(uni, basename, Scarlet)
}
