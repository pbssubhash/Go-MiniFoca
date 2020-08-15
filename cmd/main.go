package main

import (
	"flag"
	"fmt"
	"log"

	GoMiniFoca "github.com/pbssubhash/Go-MiniFoca/pkg"
)

func main() {
	fmt.Println("[+] A mini FOCA built using Golang for OSINT friends [+]")
	fmt.Println("[-] This is built for educational purposes. Author isn't responsible for misuse. [-]")
	fmt.Println(`[/\] Built by zer0 p1k4chu [/\]`)
	Target := flag.String("target", "Default", "Target website. Ex: tesla.com")
	Help := flag.Bool("help", false, "Need help?")
	Extension := flag.String("ext", "docx", "Tested values: docx")
	Pages := flag.Int("depth", 1, "Depth of Bing pages to scrape")
	Dest := flag.String("dest", "Default", "A Writable Folder to store documents")
	flag.Parse()
	if *Help == true || *Target == "Default" || *Dest == "Default" {
		flag.PrintDefaults()
		return
	}
	result, err := GoMiniFoca.Scrap(*Target, *Extension, *Pages, *Dest)
	if err != nil {
		log.Fatalf("Error with Scrap function")
	}
	for file, _ := range result {
		fmt.Println(file)
		fmt.Println(GoMiniFoca.ParseDoc(*Dest, file))
	}
}
