package goYaXml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type YandexError struct {
	YaErr YandexErrorNode `xml:"response>error"`
	err   error
}
type YandexErrorNode struct {
	Code string `xml:"code,attr"`
	Text string `xml:",chardata"`
}

type YandexResult struct {
	data     YandexSearch
	query    string
	httpCode string
	httpSize int
}

type YandexSearch struct {
	Docs []ResultYSDoc `xml:"response>results>grouping>group>doc"`
}

type ResultYSDoc struct {
	Id   string `xml:"id,attr"` // attr value
	Url  string `xml:"url"`
	Size int    `xml:"size"`
}

// YaXML – config for Query function
type YaXML struct {
	domain string // ru, com,
	user   string
	key    string
}

func Parse(xmlData []byte) (ys YandexSearch, err error) {
	err = xml.Unmarshal(xmlData, &ys)
	return
}

var YaURIprefix = "http://xmlsearch.yandex."

// Query - send query to xml yandex search, fetch & parse result
func (cfg *YaXML) Query(query string) (yr YandexResult, ye YandexError) {
	urlParts := []string{
		YaURIprefix, cfg.domain, "/xmlsearch?user=", cfg.user,
		"&key=", cfg.key, "&query=", query,
	}
	url := strings.Join(urlParts, "")
	//	fmt.Println(url)
	yr = YandexResult{YandexSearch{}, query, "", 0}

	resp, err := http.Get(url)
	if err != nil {
		ye.err = err
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ye.err = err
		return
	}

	yr.httpCode = resp.Status
	yr.httpSize = len(respData)
	if resp.Status != "200 OK" {
		ye.err = fmt.Errorf("response status is %q", resp.Status)
		return
	}

	// Check Yandex error response
	ye.err = xml.Unmarshal(respData, &ye)
	if (ye.err != nil) || (len(ye.YaErr.Code) > 0) {
		return
	}

	var ys YandexSearch
	ys, err = Parse(respData)
	if ye.err != nil {
		ye.err = err
		return
	}
	yr.data = ys

	return
}
