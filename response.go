package httpbot

import (
	"bytes"
	"io"

	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Body2Json(r io.Reader, v interface{}) error {
	bbody, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bbody, v); err != nil {
		return err
	}
	return nil
}

func Body2JsonAndString(r io.Reader, v interface{}) (string, error) {
	bbody, err := io.ReadAll(r)
	if err != nil {
		return string(bbody), err
	}
	if err = json.Unmarshal(bbody, v); err != nil {
		return string(bbody), err
	}
	return string(bbody), nil
}

func Body2String(r io.Reader) string {
	bbody, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	return string(bbody)
}

func String2Json(s string, v interface{}) error {
	if err := json.Unmarshal([]byte(s), v); err != nil {
		return err
	}
	return nil
}

func Body2Html(r io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(r)
}

func Body2HtmlAndString(r io.Reader) (*goquery.Document, string, error) {
	bbody, err := io.ReadAll(r)
	if err != nil {
		return nil, string(bbody), err
	}
	html, err := goquery.NewDocumentFromReader(bytes.NewReader(bbody))
	return html, string(bbody), err
}
