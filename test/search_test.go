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
	terms := util.SearchA(docs, "Other")
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
