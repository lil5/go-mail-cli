package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/vanng822/go-premailer/premailer"
)

var layout *template.Template

func main() {

	layout = buildTemplate()
	pagesPath := filepath.Join(".", "pages")

	err := filepath.Walk(pagesPath, func(path string, info os.FileInfo, err error) error {
		inFileName := info.Name()
		inFileExt := filepath.Ext(path)
		if info.IsDir() {
			return nil
		}
		if inFileExt != ".htm" {
			return nil
		}
		outFileName := strings.Replace(inFileName, ".gohtml", ".html", 1)

		buf := new(bytes.Buffer)
		buildPage(buf, layout, filepath.Join(pagesPath, inFileName))
		html := cssInline(buf)
		err = ioutil.WriteFile(filepath.Join(".", "build", outFileName), []byte(html), 0664)
		fmt.Print(inFileName, " -> ", outFileName, "\n")
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func buildTemplate() *template.Template {
	t, err := template.ParseGlob(filepath.Join(".", "layout", "*"))
	if err != nil {
		log.Fatal(err)
	}
	return t
}
func buildPage(buf *bytes.Buffer, t *template.Template, inFile string) {
	t, err := t.ParseGlob(filepath.Join(".", "layout", "*.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	page, err := t.ParseFiles(inFile)
	if err != nil {
		log.Fatal(err)
	}

	err = page.Execute(buf, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func cssInline(buf *bytes.Buffer) string {
	prem, err := premailer.NewPremailerFromBytes(buf.Bytes(), premailer.NewOptions())
	if err != nil {
		log.Fatal(err)
	}

	html, err := prem.Transform()
	if err != nil {
		log.Fatal(err)
	}

	return html
}
