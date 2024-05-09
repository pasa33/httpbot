package httpbot

import (
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
	if json.Unmarshal(bbody, v) != nil {
		return err
	}
	return nil
}

func Body2JsonAndString(r io.Reader, v interface{}) (string, error) {
	bbody, err := io.ReadAll(r)
	if err != nil {
		return string(bbody), err
	}
	if json.Unmarshal(bbody, v) != nil {
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

func Body2Html(r io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(r)
}
