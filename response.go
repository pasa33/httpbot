package httpbot

import (
	"encoding/json"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func Body2Json(r io.Reader, v interface{}) error {
	bbody, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if json.Unmarshal(bbody, &v) != nil {
		return err
	}
	return nil
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
