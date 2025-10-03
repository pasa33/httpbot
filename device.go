package httpbot

import (
	"maps"
	"strings"
	"sync"
)

var (
	globalDevices  = make(map[string]HeaderList)
	globalDeviceMu sync.RWMutex
)

const (
	UserAgent            headerKey = "user-agent"
	SecUa                headerKey = "sec-ch-ua"
	SecUaMobile          headerKey = "sec-ch-ua-mobile"
	SecUaArch            headerKey = "sec-ch-ua-arch"
	SecUaFullVersion     headerKey = "sec-ch-ua-full-version"
	SecUaPlatformVersion headerKey = "sec-ch-ua-platform-version"
	SecUaFullVersionList headerKey = "sec-ch-ua-full-version-list"
	SecUaBitness         headerKey = "sec-ch-ua-bitness"
	SecUaModel           headerKey = "sec-ch-ua-model"
	SecUaPlatform        headerKey = "sec-ch-ua-platform"
	AcceptLanguage       headerKey = "accept-language"
)

type headerKey string

func newHeaderKey(s string) headerKey {
	return headerKey(strings.ToLower(s))
}

type HeaderList map[headerKey]string

func (bot *HttpBot) AddDevice(id string, headers HeaderList) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)

	baseHeaders := maps.Clone(getGlobalDevice(id))
	maps.Copy(baseHeaders, headers)

	bot.devices[id] = baseHeaders
	bot.selectedDevice = baseHeaders
}

func AddGlobalDevice(id string, headers HeaderList) {
	globalDeviceMu.Lock()
	defer globalDeviceMu.Unlock()

	id = strings.ToLower(id)
	globalDevices[id] = headers
}

func getGlobalDevice(id string) HeaderList {
	globalDeviceMu.Lock()
	defer globalDeviceMu.Unlock()

	g, ok := globalDevices[id]
	if !ok {
		g = HeaderList{}
	}
	return g
}

func (bot *HttpBot) UseDevice(id string) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)
	if _, ok := bot.devices[id]; ok {
		bot.selectedDevice = bot.devices[id]
	}
}

func (h HeaderList) getValue(key headerKey) string {
	return h[key]
}
