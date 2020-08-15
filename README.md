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

##### Usage of Parameters
- Depth(depth): Number of pages to traverse inside the bing search engine
- Destination (dest): A destination folder to store the files for processesing. It should be a writable folder.
- Extension (ext): Extension to search using Bing. Currently tested on docx. Should work on xlsx and doc in theory.
- Help (help): Print Help.
- Target (target): What's the target organisation's website? Ex: tesla.com

##### For the lazy ones
```
go run main.go
./main --dest=C:\Users\User\Desktop\DownloadT\ --ext=docx --target=tesla.com --depth=2
```

### Usage as a library 

```
go get github.com/pbssubhash/Go-MiniFoca
// look at pkg/*.go for available methods.
```
