# Go-MiniFoca

### What is it? 

It's a simple metadata parser for docx files, retrieved from search engine dorks. It's a simple tool written to be a faster replacement for *some* functions of FOCA.

### Usage 

```
> go run main.go
[+] A mini FOCA built using Golang for OSINT friends [+]
[-] This is built for educational purposes. Author isn't responsible for misuse. [-]
[/\] Built by zer0 p1k4chu [/\]
  -depth int
        Depth of Bing pages to scrape (default 1)
  -dest string
        A Writable Folder to store documents (default "Default")
  -ext string
        Tested values: docx (default "docx")
  -help
        Need help?
  -target string
        Target website. Ex: tesla.com (default "Default")
```

### Usage as a library 

```
go get github.com/pbssubhash/Go-MiniFoca
// look at pkg/*.go for available methods.
```
