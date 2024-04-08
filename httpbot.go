package httpbot

import (
	"sync"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/pasa33/cookie_header"
)

type HttpBot struct {
	clientHello   profiles.ClientProfile
	clientOptions []tls_client.HttpClientOption
	client        tls_client.HttpClient
	CookieHeader  cookie_header.CookieHeader

	proxy       string
	devices     map[string]map[DeviceHeader]string
	deviceMu    sync.RWMutex
	useDeviceId string
}

func (bot *HttpBot) InitClient() error {

	options := append(bot.clientOptions, tls_client.WithProxyUrl(bot.proxy))
	options = append(options, tls_client.WithClientProfile(bot.clientHello))

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return err
	}
	bot.client = client
	return nil
}

func (bot *HttpBot) SetClientHello(ch profiles.ClientProfile) {
	bot.clientHello = ch
}

func (bot *HttpBot) SwitchClientHello(ch profiles.ClientProfile) {
	if ch.GetClientHelloStr() != bot.clientHello.GetClientHelloStr() {
		bot.SetClientHello(ch)
		bot.InitClient()
	}
}

func (bot *HttpBot) SetClientOptions(options []tls_client.HttpClientOption) {
	bot.clientOptions = options
}

// call before init client
func (bot *HttpBot) InitProxy(proxyUrl string) {
	if bot.proxy == "" {
		bot.proxy = proxyUrl
	}
}

// call when client already inuse
func (bot *HttpBot) SwitchProxy(proxyUrl string) {
	bot.proxy = proxyUrl
	bot.client.SetProxy(bot.proxy)
}
