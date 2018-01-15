package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

const targetUrl = "http://www.apodro.ch/"
const urlProduct = "https://www.xtrapharm.ch"

// Start search here (dig x levels deeper from here to find the a element)
const elRegex = `<div class="daydeal`

// $1 in this regex will be the thing we are interested in:
const prodRegex = `<a href="https://www.xtrapharm.ch/(.*)\.html`

// This requests the "daydeal" of Apodro.ch to have interesting products on the radar automatically. Beware that Apodro can change the DOM structure at any time by their free will ;-)
func main() {
	resp, err := http.Get(targetUrl)
	if err != nil {
		log.Fatalln("ERROR: Could not read from", targetUrl, err)
	}
	defer resp.Body.Close()
	var s []byte
	resp.Body.Read(s)

	doc, err := html.Parse(resp.Body)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					// fmt.Println(a.Val)
					hasDaydealInSrc := strings.Contains(a.Val, "daydeal")
					if hasDaydealInSrc {
						u := getHrefUrl(n.Parent) // get href for the parent <a>
						fmt.Println(u.String())
						// fmt.Println(targetUrl + a.Val)
					}
					break // break after the first (hopefully only) href
				}
			}
		}
		// Recursion here
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func getProductName(url string) string {
	return ""
}

func getHrefUrl(n *html.Node) (u *url.URL) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				hasProto := strings.Index(a.Val, "http") == 0
				if hasProto {
					u, err := url.Parse(a.Val)
					if err != nil {
						log.Fatal(err)
					}
					return u
				}
			}
		}
	}
	return
}

func getProductImg(url string) image.Image {
	return nil
}
