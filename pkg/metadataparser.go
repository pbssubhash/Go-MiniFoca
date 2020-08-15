package GoMiniFoca

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type CoreProperties struct {
	XMLName        xml.Name `xml:"coreProperties"`
	Title          string   `xml:"title"`
	Creator        string   `xml:"creator"`
	Keywords       string   `xml:"keywords"`
	Description    string   `xml:"description"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
	Revision       string   `xml:"revision"`
	Created        struct {
		Text string `xml:",chardata"`
	} `xml:"created"`
	Modified struct {
		Text string `xml:",chardata"`
	} `xml:"modified"`
}

type Properties struct {
	XMLName      xml.Name `xml:"Properties"`
	Template     string   `xml:"Template"`
	Pages        string   `xml:"Pages"`
	Words        string   `xml:"Words"`
	Characters   string   `xml:"Characters"`
	Lines        string   `xml:"Lines"`
	Application  string   `xml:"Application"`
	HeadingPairs struct {
		Vector struct {
			Variant []struct {
				Lpstr string `xml:"lpstr"`
			} `xml:"variant"`
		} `xml:"vector"`
	} `xml:"HeadingPairs"`
	TitlesOfParts struct {
		Vector struct {
			Lpstr string `xml:"lpstr"`
		} `xml:"vector"`
	} `xml:"TitlesOfParts"`
	Company    string `xml:"Company"`
	AppVersion string `xml:"AppVersion"`
}

func checkError(ok error) {
	if ok != nil {
		log.Fatalf("Error occured.")
	}
}

// Taken from https://stackoverflow.com/a/58192644
func Unzip(src, dest string) error {
	dest = filepath.Clean(dest) + string(os.PathSeparator)
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()
	os.MkdirAll(dest, 0755)
	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		path := filepath.Join(dest, f.Name)
		// Check for ZipSlip: https://snyk.io/research/zip-slip-vulnerability
		if !strings.HasPrefix(path, dest) {
			return fmt.Errorf("%s: illegal file path", path)
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
func ParseXML(file string, types string) (map[string]string, error) {
	fmt.Println(file)
	xmlReader, ok := os.Open(file)
	if ok != nil {
		log.Fatalf("Error-1")
	}
	content, _ := ioutil.ReadAll(xmlReader)
	// fmt.Println(string(content))
	switch types {
	case "app":
		var App *Properties
		xml.Unmarshal([]byte(content), &App)
		return map[string]string{
			"Title":       App.HeadingPairs.Vector.Variant[0].Lpstr,
			"Application": fmt.Sprintf("%s %s", App.Application, App.AppVersion),
			"Words":       App.Words,
			"Company":     App.Company,
			"Pages":       App.Pages,
			"Template":    App.Template}, nil
	case "core":
		var App *CoreProperties
		xml.Unmarshal([]byte(content), &App)
		return map[string]string{
			"CreatedBy":    App.Creator,
			"CreatedTime":  App.Created.Text,
			"ModifiedTime": App.Modified.Text,
			"Description":  App.Description,
			"Keywords":     App.Keywords}, nil
	default:
		return nil, nil
	}
}

func ParseDoc(DestFolder string, ZipFile string) (map[string]string, map[string]string) {
	// fmt.Println(ZipFile)
	// fmt.Println(DestFolder)
	// fmt.Println(DestFolder + strings.TrimSuffix(path.Base(ZipFile), path.Ext(path.Base(ZipFile))))
	// fmt.Println(strings.TrimSuffix(path.Base(ZipFile), path.Ext(path.Base(ZipFile))))
	FullFileLoc := strings.TrimSuffix(ZipFile, path.Ext(ZipFile))
	fmt.Println(FullFileLoc)
	Unzip(ZipFile, FullFileLoc)
	appmap, ok := ParseXML(fmt.Sprintf("%s%sdocProps%score.xml", FullFileLoc, string(os.PathSeparator), string(os.PathSeparator)), "core")
	if ok != nil {
		log.Fatalf("Error - 2")
	}
	coremap, ok := ParseXML(fmt.Sprintf("%s%sdocProps%sapp.xml", FullFileLoc, string(os.PathSeparator), string(os.PathSeparator)), "app")
	if ok != nil {
		log.Fatalf("Error - 3")
	}
	return appmap, coremap
}
