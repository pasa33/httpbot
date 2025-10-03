package httpbot

import (
	"sync"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/pasa33/cookie_header"
)

type HttpBot struct {
	client        tls_client.HttpClient
	clientHello   profiles.ClientProfile
	clientOptions []tls_client.HttpClientOption
	proxy         string
	CookieHeader  cookie_header.CookieHeader
	//devices
	devices        map[string]HeaderList
	selectedDevice HeaderList
	deviceMu       sync.RWMutex
	//flags
	skipEmptyHeaders bool
	isDebug          bool
	isInited         bool
}

func New() *HttpBot {
	return &HttpBot{
		devices:        map[string]HeaderList{},
		selectedDevice: HeaderList{},
		isInited:       false,
	}
}

func (bot *HttpBot) InitClient() (err error) {
	options := append(bot.clientOptions, tls_client.WithProxyUrl(bot.proxy))
	options = append(options, tls_client.WithClientProfile(bot.clientHello))

	bot.client, err = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err == nil {
		bot.isInited = true
	}
	return
}

func (bot *HttpBot) SetClientHello(ch profiles.ClientProfile) {
	if ch.GetClientHelloStr() != bot.clientHello.GetClientHelloStr() {
		return
	}
	bot.clientHello = ch
	if bot.isInited {
		bot.InitClient()
	}
}

func (bot *HttpBot) SetClientOptions(options []tls_client.HttpClientOption) {
	bot.clientOptions = options
	if bot.isInited {
		bot.InitClient()
	}
}

// call for first time before client inuse
func (bot *HttpBot) SetProxy(proxyUrl string) {
	if bot.proxy == proxyUrl {
		return
	}
	bot.proxy = proxyUrl
	if bot.isInited {
		bot.client.SetProxy(bot.proxy)
	}
}

func (bot *HttpBot) InitCookieHeader() {
	bot.CookieHeader = cookie_header.New()
}

func (bot *HttpBot) SetSkipEmptyHeaders(skip bool) {
	bot.skipEmptyHeaders = skip
}

func (bot *HttpBot) SetDebug(b bool) {
	bot.isDebug = b
}
