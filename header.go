package httpbot

import (
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

type Header struct {
	Key   string
	Value string
	//SkipIfEmpty bool
}

const (
	Auto_Header = "!#!AutoHeader!#!"
)

func (bot *HttpBot) generateHeaders(headers []Header) map[string][]string {
	hds := make(map[string][]string)
	for _, v := range headers {
		if bot.skipEmptyHeaders && v.Value == "" {
			continue
		}
		//append to header-order
		hds[http.HeaderOrderKey] = append(hds[http.HeaderOrderKey], strings.ToLower(v.Key))
		//add header value
		if v.Value == Auto_Header {
			if isOnlyOrder(v.Key) {
				continue
			}
			hds[v.Key] = []string{bot.getInUseDeviceValue(v.Key)}
		} else {
			hds[v.Key] = []string{v.Value}
		}
	}
	return hds
}

func isOnlyOrder(s string) bool {
	switch strings.ToLower(s) {
	case "content-length":
		return true
	}
	return false
}
