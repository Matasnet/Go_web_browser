package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var url string

	if len(os.Args) < 2 {
		fmt.Print("Podaj adres strony internetowej: ")
		fmt.Scan(&url)
	} else {
		url = os.Args[1]
	}

	// Dodaj protokół "https://" w przypadku braku.
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Błąd podczas pobierania strony:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Błąd podczas analizy HTML:", err)
		os.Exit(1)
	}

	printTextNodes(doc)
}

func printTextNodes(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			printTextNodes(c)
		}
	}
}
