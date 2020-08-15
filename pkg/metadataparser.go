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

func Unzip(src string, destination string) ([]string, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(destination, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return nil, fmt.Errorf("%s is an illegal filepath", fpath)
		}
		// filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			fmt.Println(err.Error())
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fmt.Println(err.Error())
		}
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		return nil, nil
	}
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
	FullFileLoc := strings.Split(strings.TrimSuffix(ZipFile, path.Ext(ZipFile)), ".")[0]
	fmt.Println(FullFileLoc)
	_, ok := Unzip(ZipFile, FullFileLoc)
	if ok != nil {
		fmt.Println(string(ok.Error()))
	}
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
