package goyaxml

import "io/ioutil"
import "testing"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestParse(t *testing.T) {
	var tests = []struct {
		testXMLfile string
		want        YandexSearch
		docsCount   int
	}{
		{"ya_response_01.xml", YandexSearch{}, 10},
	}
	for _, c := range tests {
		xmlData, err := ioutil.ReadFile(c.testXMLfile)
		check(err)
		got, _ := Parse(xmlData)
		gotDocsCount := len(got.Docs)
		if gotDocsCount != c.docsCount {
			t.Errorf("Documents count is %d, but must be %d", gotDocsCount, c.docsCount)
		}
	}
}

// developer's test, skip it
func _TestQuery(t *testing.T) {
	yaFetch := YaXML{"com", "your_login_here", "your_key_here"}
	yr, _ := yaFetch.Query("sex")
	gotDocsCount := len(yr.data.Docs)
	if gotDocsCount != 10 {
		t.Errorf("Fetch documents count is %d, but must be 10", gotDocsCount)
	}
}
