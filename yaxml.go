package yaXML

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type YandexError struct {
	YaErr YandexErrorNode `xml:"response>error"`
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

// CfgYaXML – config for Query function
type CfgYaXML struct {
	domain string // ru, com,
	user   string
	key    string
}

func Parse(xmlData []byte) (ys YandexSearch, err error) {
	err = xml.Unmarshal(xmlData, &ys)
	return
}

func Config(domain, user, key string) (cfg CfgYaXML) {
	cfg.domain = domain
	cfg.user = user
	cfg.key = key
	return
}

var YaURIprefix = "http://xmlsearch.yandex."

// Query - send query to xml yandex search, fetch & parse result
func (cfg *CfgYaXML) Query(query string) (yr YandexResult, err error) {
	var ye YandexError
	urlParts := []string{
		YaURIprefix, cfg.domain, "/xmlsearch?user=", cfg.user,
		"&key=", cfg.key, "&query=", query,
	}
	url := strings.Join(urlParts, "")
	//	fmt.Println(url)
	yr = YandexResult{YandexSearch{}, query, "", 0}

	resp, err := http.Get(url)
	if err != nil {
		//ye.err = err
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//ye.err = err
		return
	}

	yr.httpCode = resp.Status
	yr.httpSize = len(respData)
	if resp.Status != "200 OK" {
		err = fmt.Errorf("response status is %s", resp.Status)
		return
	}

	// Check Yandex error response
	err = xml.Unmarshal(respData, &ye)
	if err != nil {
		return
	}
	if len(ye.YaErr.Code) > 0 {
		err = fmt.Errorf("Yandex.XML returned code: %s, text: %s",
			ye.YaErr.Code, ye.YaErr.Text)
		return
	}

	var ys YandexSearch
	ys, err = Parse(respData)
	if err != nil {
		return
	}
	yr.data = ys

	return
}
