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
	Target := flag.String("t", "Default", "Target website. Ex: tesla.com")
	Help := flag.Bool("h", false, "Need help?")
	Extension := flag.String("e", "docx", "Tested values: docx")
	Pages := flag.Int("p", 1, "Depth of Bing pages to scrape")
	Dest := flag.String("d", "Default", "A Writable Folder to store documents")
	if *Help == true || *Target == "Default" || *Dest == "Default" {
		flag.PrintDefaults()
	}
	result, err := GoMiniFoca.Scrap(*Target, *Extension, *Pages, *Dest)
	if err != nil {
		log.Fatalf("Error with Scrap function")
	}
}
