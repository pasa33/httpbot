package httpbot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

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
	res, err := bot.client.Do(req)
	if err != nil {
		log.Printf("%+v", err)
	}

	return res, err
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

func EncodeJSON(j map[string]interface{}) []byte {
	jsonData, _ := json.Marshal(j)
	return jsonData
}

func EncodeURLForm(j map[string]interface{}) []byte {
	if len(j) == 0 {
		return []byte{}
	}
	var buf strings.Builder
	for k, v := range j {
		buf.WriteByte('&')
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(fmt.Sprint(v))
	}
	return []byte(buf.String()[1:])
}

func DecompressBody(res *http.Response) io.ReadCloser {
	return http.DecompressBody(res)
}
