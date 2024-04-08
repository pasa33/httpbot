package httpbot

import (
	"bytes"
	"io"

	http "github.com/bogdanfinn/fhttp"
)

func (bot *HttpBot) MakeRequest(method, url string, headers []Header, payload ...[]byte) (red *http.Response, err error) {

	var body io.Reader = nil
	if len(payload) > 0 && len(payload[0]) > 0 {
		body = bytes.NewBuffer(payload[0])
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = bot.generateHeaders(headers)

	return bot.client.Do(req)
}

func (bot *HttpBot) MakeRequestCustomOrder(method, url string, headers []Header, hOrder, pOrder []string, payload ...[]byte) (red *http.Response, err error) {

	var body io.Reader = nil
	if len(payload) > 0 && len(payload[0]) > 0 {
		body = bytes.NewBuffer(payload[0])
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = bot.generateHeaders(headers)

	if len(hOrder) > 0 {
		req.Header[http.HeaderOrderKey] = hOrder
	}

	if len(pOrder) > 0 {
		req.Header[http.PHeaderOrderKey] = pOrder
	}

	return bot.client.Do(req)
}
