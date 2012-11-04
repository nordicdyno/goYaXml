goYaXml
=======

About
--------

`goYaXml` – Go lib for fetch & parse Yandex XML respones

<quote>Yandex.XML is a service that lets you send queries to the Yandex search engine
and get responses in XML format.</quote>

Official Yandex.XML search (need login):
* Russian: http://xml.yandex.ru/settings.xml
* English: http://xml.yandex.com/settings.xml

Warning
--------

This lib is just my game with Go language, it's usable but not completely parser
(needs a bit more work for handling whole response)

And also it just support only `Get` requests

Installation
--------

go get github.com/nordicdyno/goYaXml


Usage
---------

    import "fmt"
    import "github.com/nordicdyno/goYaXml"
    
    yaFetchCfg := yaXML.Config("com", "your_login_here", "your_key_here")
    yResultsXML, err := yaFetchCfg.Query("safe_uri_escaped_query_string_here")
    if err != nil {
        panic(err)
    }
    fmt.Println(yResultsXML)



