package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"testing"
	"text/template"

	"github.com/Jun-Chang/vast-xml-test/qtpl"
	"github.com/Jun-Chang/vast-xml-test/vast"
)

var t = template.Must(template.ParseFiles("./data/vast_template.xml"))

func BenchmarkXMLMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := vast.New()
		xml.Marshal(v)
	}
}

func BenchmarkXMLTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := vast.New()
		var b bytes.Buffer

		w := bufio.NewWriter(&b)

		t.Execute(w, v)
		w.Flush()
		b.Bytes()
	}
}

func BenchmarkXMLQuickTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := vast.New()
		_ = []byte(qtpl.XMLFromQuickTemplate(v))
	}
}
