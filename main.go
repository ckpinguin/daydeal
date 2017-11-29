package main

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "http://www.apodro.ch/"
const urlProduct = "https://www.xtrapharm.ch"

// Start search here (dig x levels deeper from here to find the a element)
const elRegex = `<div class="daydeal`

// $1 in this regex will be the thing we are interested in:
const prodRegex = `<a href="https://www.xtrapharm.ch/(.*)\.html`

// This requests the "daydeal" of Apodro.ch to have interesting products on the radar automatically. Beware that Apodro can change the DOM structure at any time by their free will ;-)
func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(fmt.Sprint("Cannot read URL", url, "Error:", err.Error()))
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(fmt.Sprint("Cannot read response body:", err.Error()))
		os.Exit(1)
		res.Body.Close()
	}
	defer res.Body.Close()

	r := bytes.NewReader(b)
	_ = r
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
