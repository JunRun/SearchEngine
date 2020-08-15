package util

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	s := Tokenize("This is A BIG dog")
	l := LowerFilter(s)
	term := StopWordsFilter(l)
	fmt.Println(term)
}

func TestSelect(t *testing.T) {
	go func() {
		for i := 0; i >= 0; i++ {
			fmt.Println(i)
		}
	}()
	select {}
	fmt.Println("12312311")
}
