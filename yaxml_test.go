package yaXML

import "io/ioutil"
import "testing"
import "os"

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

func TestQuery(t *testing.T) {
	// developer's test, skip it
	if len(os.Getenv("DEVELOPER_TESTS")) == 0 {
		return
	}
	yaFetch := Config("com", os.Getenv("YA_XML_LOGIN"), os.Getenv("YA_XML_KEY"))
	yr, err := yaFetch.Query("sex")
	if err != nil {
		t.Errorf("Error %s", err)
	}
	gotDocsCount := len(yr.data.Docs)
	if gotDocsCount != 10 {
		t.Errorf("Fetch documents count is %d, but must be 10", gotDocsCount)
	}
}
