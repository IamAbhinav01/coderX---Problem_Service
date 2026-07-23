package utils

import (
	"bytes"
	"log"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var mdParser = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)

var policy = bluemonday.UGCPolicy()

func Sanitize_to_HTML(markdown string) (string, error) {
	var buf bytes.Buffer

	err:= mdParser.Convert([]byte(markdown),&buf)
	if err != nil{
		log.Println("Error occured while santiiling the markdown")
		return "",err
	}
	safeHtml:=policy.SanitizeBytes(buf.Bytes())

	sanitizedMarkdown, err := htmltomarkdown.ConvertString(string(safeHtml))
	if err != nil {
		log.Println("Error occurred while converting HTML back to Markdown:", err)
		return "", err
	}

	return sanitizedMarkdown, nil
}