package util

import (
	"fmt"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"text"`
	ID    int
}

//加载 xml 文件资料
func loadDocument(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("the path was error")
		return nil, err
	}
	defer f.Close()
	return nil, err

}
