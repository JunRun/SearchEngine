package test

import (
	"engine/util"
	"fmt"
	"testing"
)

func TestSearchA(t *testing.T) {

	docs, err := util.LoadDocument("../resource/enwiki-latest-abstract1.xml")
	if err != nil {
		return
	}
	terms := util.SearchA(docs, "European settlement")
	for _, term := range terms {
		fmt.Println(term.Title, term.ID)
	}
}

func TestSearchB(t *testing.T) {

	docs, err := util.LoadDocument("../resource/enwiki-latest-abstract1.xml")
	if err != nil {
		return
	}
	terms := util.SearchB(docs, "`(?i)\bother\b`")
	for _, term := range terms {
		fmt.Println(term.Title, term.ID)
	}
}
func TestInvertSearch(t *testing.T) {
	docs, err := util.LoadDocument("../resource/enwiki-latest-abstract1.xml")
	if err != nil {
		return
	}
	idx := make(util.Index)
	idx.Add(docs)
	fmt.Println(idx.InvertSearch("European settlement"))

}
func TestAdd(t *testing.T) {
	idx := make(util.Index)
	idx.Add([]util.Document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	idx.Add([]util.Document{{ID: 2, Text: "donut is a donut"}})
	fmt.Println(idx)
}
