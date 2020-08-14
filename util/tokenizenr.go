package util

import (
	snowballeng "github.com/kljensen/snowball/english"
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {

		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

//将单词全部变为小写
func LowerFilter(token []string) []string {
	r := make([]string, len(token))
	for i, t := range token {
		r[i] = strings.ToLower(t)
	}
	return r
}

var stopWords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

//过滤字符串
func StopWordsFilter(token []string) []string {
	r := make([]string, 0, len(token))
	for _, t := range token {
		if _, ok := stopWords[t]; !ok {
			r = append(r, t)
		}
	}
	return r
}

//词根化
func StemmerFilter(token []string) []string {
	r := make([]string, len(token))
	for i, t := range token {
		r[i] = snowballeng.Stem(t, false)
	}
	return r
}
func Filter(text string) []string {
	s := Tokenize(text)
	l := LowerFilter(s)
	w := StopWordsFilter(l)
	term := StemmerFilter(w)
	return term
}
