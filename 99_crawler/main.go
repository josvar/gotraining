package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	url := "http://www.lanacion.com.ar/sitemap-news.xml"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
