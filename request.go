package httpbot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (bot *HttpBot) PrepareRequest(method, url string, headers []Header, payload ...[]byte) (reqq *http.Request, err error) {

	var body io.Reader = nil
	if len(payload) > 0 && len(payload[0]) > 0 {
		body = bytes.NewBuffer(payload[0])
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = bot.generateHeaders(headers)

	return req, err
}

func (bot *HttpBot) SendRequest(req *http.Request) (red *http.Response, err error) {

	res, err := bot.client.Do(req)
	if err != nil {
		err, parsed := parseRequestError(err)
		if !parsed {
			log.Printf("%+v", err)
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

	var body io.Reader = nil
	if len(payload) > 0 && len(payload[0]) > 0 {
		body = bytes.NewBuffer(payload[0])
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, nil, err
	}

	req.Header = bot.generateHeaders(headers)
	res, err := bot.client.Do(req)
	if err != nil {
		err, parsed := parseRequestError(err)
		if !parsed {
			log.Printf("%+v", err)
		}
	}

	return req, res, err
}

func EncodeJSON(j map[string]any) []byte {
	jsonData, _ := json.Marshal(j)
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

func parseRequestError(err error) (error, bool) {
	str := strings.ToLower(err.Error())
	if strings.Contains(str, `client.timeout exceeded while awaiting headers`) {
		return fmt.Errorf("timeout exceeded"), true
	}

	if strings.Contains(str, `proxy responded with non 200 code:`) {
		code := strings.Split(str, "proxy responded with non 200 code:")[1]
		return fmt.Errorf("proxy error: %s", code), true
	}

	return err, false
}
