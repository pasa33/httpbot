package httpbot

import (
	"strings"
)

type deviceHeader string

const (
	UserAgent            deviceHeader = "user-agent"
	SecUa                deviceHeader = "sec-ch-ua"
	SecUaMobile          deviceHeader = "sec-ch-ua-mobile"
	SecUaArch            deviceHeader = "sec-ch-ua-arch"
	SecUaFullVersion     deviceHeader = "sec-ch-ua-full-version"
	SecUaPlatformVersion deviceHeader = "sec-ch-ua-platform-version"
	SecUaFullVersionList deviceHeader = "sec-ch-ua-full-version-list"
	SecUaBitness         deviceHeader = "sec-ch-ua-bitness"
	SecUaModel           deviceHeader = "sec-ch-ua-model"
	SecUaPlatform        deviceHeader = "sec-ch-ua-platform"
	AcceptLanguage       deviceHeader = "accept-language"
)

func (bot *HttpBot) AddDevice(id string, headers map[deviceHeader]string) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)
	bot.devices[id] = headers
	if len(bot.devices) == 1 {
		bot.UseDevice(id)
	}
}

func (bot *HttpBot) EditDevice(id string, headers map[deviceHeader]string) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)
	if _, ok := bot.devices[id]; !ok {
		return
	}
	for k, v := range headers {
		bot.devices[id][k] = v
	}
}

func (bot *HttpBot) UseDevice(id string) {
	bot.deviceMu.RLock()
	defer bot.deviceMu.RUnlock()

	id = strings.ToLower(id)
	if _, ok := bot.devices[id]; ok {
		bot.useDeviceId = id
	}
}

func (bot *HttpBot) getInUseDeviceValue(key string) string {
	bot.deviceMu.RLock()
	defer bot.deviceMu.RUnlock()

	if val, ok := bot.devices[bot.useDeviceId]; ok {
		key = strings.ToLower(key)
		return val[deviceHeader(key)]
	}
	return ""
}
