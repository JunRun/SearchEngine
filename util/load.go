package util

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

//加载 xml 文件资料
func LoadDocument(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("the path is error")
		return nil, err
	}
	defer f.Close()
	//将文件加载到内存中
	content := xml.NewDecoder(f)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := content.Decode(&dump); err != nil {
		fmt.Println("the decode was error")
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, err

}
func SearchA(docs []Document, term string) []Document {
	var r []Document
	for _, doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}
