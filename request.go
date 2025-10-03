package httpbot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bytedance/sonic"
)

func (bot *HttpBot) PrepareRequest(method, url string, headers []Header, payload ...[]byte) (reqq *http.Request, err error) {

	var body io.Reader = nil
	if len(payload) > 0 && len(payload[0]) > 0 {
		body = bytes.NewBuffer(payload[0])
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		if bot.isDebug {
			log.Printf("[%s] %+v", bot.proxy, err)
		}
		return nil, err
	}

	req.Header = bot.generateHeaders(headers)

	return req, err
}

func (bot *HttpBot) SendRequest(req *http.Request) (red *http.Response, err error) {

	res, err := bot.client.Do(req)
	if err != nil {
		if bot.isDebug {
			log.Printf("[%s] %+v", bot.proxy, err)
		}

		if bot.proxy != "" {
			bot.SetProxy(bot.proxy) //ricreo il transport
		}
	}

	return res, err
}

func (bot *HttpBot) MakeRequest(method, url string, headers []Header, payload ...[]byte) (red *http.Response, err error) {

	req, err := bot.PrepareRequest(method, url, headers, payload...)
	if err != nil {
		return nil, err
	}
	return bot.SendRequest(req)
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

func (bot *HttpBot) MakeReturnRequest(method, url string, headers []Header, payload ...[]byte) (reqq *http.Request, ress *http.Response, err error) {

	req, err := bot.PrepareRequest(method, url, headers, payload...)
	if err != nil {
		return req, nil, err
	}

	res, err := bot.SendRequest(req)
	return req, res, err
}

func EncodeJSON(j map[string]any) []byte {
	jsonData, _ := sonic.Marshal(j)
	return jsonData
}

func EncodeURLForm(j map[string]any) []byte {
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
