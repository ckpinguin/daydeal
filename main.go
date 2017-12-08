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

const URL = "http://www.apodro.ch/"
const urlProduct = "https://www.xtrapharm.ch"

// Start search here (dig x levels deeper from here to find the a element)
const elRegex = `<div class="daydeal`

// $1 in this regex will be the thing we are interested in:
const prodRegex = `<a href="https://www.xtrapharm.ch/(.*)\.html`

// This requests the "daydeal" of Apodro.ch to have interesting products on the radar automatically. Beware that Apodro can change the DOM structure at any time by their free will ;-)
func main() {
	// res, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalln(fmt.Sprint("Cannot read URL", url, "Error:", err.Error()))
	// 	os.Exit(1)
	// }
	// b, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(fmt.Sprint("Cannot read response body:", err.Error()))
	// 	os.Exit(1)
	// 	res.Body.Close()
	// }
	// defer res.Body.Close()

	// r := bytes.NewReader(b)
	// _ = r

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln("ERROR: Could not read from", URL, err)
	}
	defer resp.Body.Close()
	var s []byte
	resp.Body.Read(s)

	doc, err := html.Parse(resp.Body)
	var f func(*html.Node)
	f = func(n *html.Node) {
		// fmt.Println("Starting parse...")
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// fmt.Println(a.Val)
					hasProto := strings.Index(a.Val, "http") == 0
					if hasProto {
						u, err := url.Parse(a.Val)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(u.String())
					}
					break // break after the first (hopefully only) href
				}
			}
		}
		// Recurse
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// root, _ := html.Parse(r)
	// log.Println(root)
	// div := htmlnode.Find(root, `<div class="daydeal view view-daydeal view-id-daydeal view-display-id-block_1 js-view-dom-id-*">`)[0]
	// anchor := htmlnode.Find(div, `<a>`)
	// log.Println(anchor)
	// log.Println(anchor[0].Attr)
}

func getProductName(url string) string {
	return ""
}

func getProductImg(url string) image.Image {
	return nil
}
